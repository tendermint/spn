import json
import subprocess
import datetime
import time

valAddress = 'spnvaloper1ezptsm3npn54qx9vvpah4nymre59ykr9mx22g4'
saleTemplateFile = './sale_template.json'
saleFile = './sale.json'

def cmd(command):
    subprocess.run([command], shell=True, check=True)

def cmd_devnull(command):
    subprocess.run([command], shell=True, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)

def date_f(d):
    return d.isoformat("T") + "Z"

def set_sale_json(sellingDenom, sellingAmount, payingDenom, price, startTime, endTime):
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

# Initialization
cmd_devnull('spnd tx staking delegate {} 100000uspn --from bob -y'.format(valAddress))
cmd_devnull('spnd tx staking delegate {} 100000uspn --from carol -y'.format(valAddress))
cmd_devnull('spnd tx staking delegate {} 100000uspn --from dave -y'.format(valAddress))
cmd_devnull('spnd tx profile create-coordinator --from alice -y')
cmd_devnull('spnd tx campaign create-campaign orbit 1000000orbit --from alice -y')
cmd_devnull('spnd tx campaign mint-vouchers 1 100000orbit --from alice -y')

# Define auction start and end from current time
dateNow = datetime.datetime.utcnow()
start = dateNow + datetime.timedelta(seconds=15)
end = dateNow + datetime.timedelta(seconds=10000)

# Fundraising
set_sale_json('v/1/orbit', '50000', 'uspn', '100', date_f(start), date_f(end))
cmd_devnull('spnd tx fundraising create-fixed-price-auction {} --from alice -y'.format(saleFile))
cmd_devnull('spnd tx participation participate 1 4 --from bob -y')
cmd_devnull('spnd tx participation participate 1 4 --from carol -y')
cmd_devnull('spnd tx participation participate 1 4 --from dave -y')

# Wait auction start
print("waiting for auction start...")
time.sleep(15)

# Place bid
cmd('spnd tx fundraising bid 1 fixed-price 100 10000v/1/orbit --from bob -y')
cmd('spnd tx fundraising bid 1 fixed-price 100 20000v/1/orbit --from carol -y')
cmd('spnd tx fundraising bid 1 fixed-price 100 20000v/1/orbit --from dave -y')

# Wait withdrawal delay
print("waiting for withdrawal delay...")
time.sleep(5)

cmd_devnull('spnd tx participation withdraw-allocations 1 --from bob -y')
cmd_devnull('spnd tx participation withdraw-allocations 1 --from carol -y')
cmd_devnull('spnd tx participation withdraw-allocations 1 --from dave -y')