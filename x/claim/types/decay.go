package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewDisabledDecay returns decay information with disabled decay
func NewDisabledDecay() DecayInformation {
	// convert all time to UTC
	return DecayInformation{
		Enabled:    false,
		DecayStart: time.UnixMilli(0).UTC(),
		DecayEnd:   time.UnixMilli(0).UTC(),
	}
}

// NewEnabledDecay returns decay information with a decay start and end
func NewEnabledDecay(start, end time.Time) DecayInformation {
	// convert all time to UTC
	return DecayInformation{
		Enabled:    true,
		DecayStart: start.UTC(),
		DecayEnd:   end.UTC(),
	}
}

// Validate validates the decay information
func (m DecayInformation) Validate() error {
	if m.Enabled && m.DecayStart.After(m.DecayEnd) {
		return fmt.Errorf("decay starts after decay end %s > %s", m.DecayStart.String(), m.DecayEnd.String())
	}

	return nil
}

// ApplyDecayFactor reduces the coins depending on the decay factor from decay information
// coins decrease from decay start to zero at decay end
func (m DecayInformation) ApplyDecayFactor(coins sdk.Coins, currentTime time.Time) sdk.Coins {
	// no decay factor applied
	if coins.Empty() || !m.Enabled || currentTime.Before(m.DecayStart) {
		return coins
	}

	// coins reduced to 0 if decay ended
	if currentTime.Equal(m.DecayEnd) || currentTime.After(m.DecayEnd) {
		return sdk.NewCoins()
	}

	// calculate decay factor
	timeToDec := func(t time.Time) sdk.Dec {
		return sdk.NewDecFromInt(sdkmath.NewInt(t.Unix()))
	}

	current, start, end := timeToDec(currentTime), timeToDec(m.DecayStart), timeToDec(m.DecayEnd)

	// (end-current)/(end-start)
	decayFactor := (end.Sub(current)).Quo(end.Sub(start))

	// apply decay factor to each denom
	newCoins := sdk.NewCoins()
	for _, coin := range coins {
		amountDec := sdk.NewDecFromInt(coin.Amount)
		newAmount := amountDec.Mul(decayFactor).TruncateInt()

		if !newAmount.IsZero() {
			newCoins = append(newCoins, sdk.NewCoin(
				coin.Denom,
				newAmount,
			))
		}
	}

	return newCoins.Sort()
}
