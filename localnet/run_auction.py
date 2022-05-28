import json
import os
import datetime
from utils import cmd_devnull, cmd, initialize_campaign, date_f

auction_template_file = './auctions/auction_template.json'
auction_file = './auctions//auction.json'

def set_sale_json(selling_denom, selling_amount, paying_denom, start_price, min_bid_price, start_time, end_time):
    f = open(auction_template_file)
    jf = json.load(f)
    jf['selling_coin']['denom'] = selling_denom
    jf['selling_coin']['amount'] = selling_amount
    jf['paying_coin_denom'] = paying_denom
    jf['start_price'] = start_price
    jf['min_bid_price'] = min_bid_price
    jf['start_time'] = start_time
    jf['end_time'].append(end_time)
    with open(auction_file, 'w', encoding='utf-8') as newF:
        json.dump(jf, newF, ensure_ascii=False, indent=4)

if __name__ == "__main__":
    initialize_campaign()

    # Define auction start and end from current time
    date_now = datetime.datetime.utcnow()
    start = date_now + datetime.timedelta(seconds=15)
    end = date_now + datetime.timedelta(seconds=40)

    # Fundraising
    set_sale_json('v/1/orbit', '50000', 'uspn', '100', '50', date_f(start), date_f(end))
    cmd_devnull('spnd tx fundraising create-fixed-price-auction {} --from alice -y'.format(auction_file))
    os.remove(auction_file)
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