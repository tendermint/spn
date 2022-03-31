import json
import subprocess
import datetime

saleTemplateFile = './saleTemplate.json'
saleFile = './sale.json'

def setSale(sellingDenom, sellingAmount, payingDenom, price, startTime, endTime):
    f = open(saleTemplateFile)
    jf = json.load(f)
    jf['selling_coin']['denom'] = sellingDenom
    jf['selling_coin']['amount'] = sellingAmount
    jf['paying_coin_denom'] = payingDenom
    jf['start_price'] = price
    jf['start_time'] = startTime
    jf['end_time'] = endTime
    with open(saleFile, 'w', encoding='utf-8') as newF:
        json.dump(jf, newF, ensure_ascii=False, indent=4)

def datef(d):
    return d.isoformat("T") + "Z"

dateNow = datetime.datetime.utcnow()
start = dateNow + datetime.timedelta(seconds=15)
end = dateNow + datetime.timedelta(seconds=60)

setSale('v/1/orbit', '10000', 'uspn', '100', datef(start), datef(end))

# Initialization
subprocess.run(['spnd tx staking delegate spnvaloper1ezptsm3npn54qx9vvpah4nymre59ykr9mx22g4 100000uspn --from bob -y'], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)

subprocess.run(['spnd tx profile create-coordinator --from alice -y'], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
subprocess.run(['spnd tx campaign create-campaign orbit 1000000orbit --from alice -y'], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
subprocess.run(['spnd tx campaign mint-vouchers 1 50000orbit --from alice -y'], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)

subprocess.run(['spnd tx fundraising create-fixed-price-auction sale.json --from alice -y'], shell=True, check=True)
subprocess.run(['spnd tx participation participate 1 2 --from bob -y'], shell=True, check=True)
# spnd tx participation withdraw-allocations 1 --from bob -y