# SPN_NODE env var should contain the node to connect to
if [[ -z "${SPN_NODE}" ]]; then
  SPN_NODE="http://0.0.0.0:26657"
else
  SPN_NODE="${SPN_NODE}"
fi

# check spnd is installed
spnd > /dev/null
if [ $? != 0 ]; then
   echo "spnd binary not installed"
   exit 1
fi

CreateAccount () {
  spnd keys delete $1 -y > /dev/null
  echo $2 | spnd keys add $1 --recover > /dev/null
}

SpndCommand () {
  command="spnd tx $1 --from $2 -y --node ${SPN_NODE}"
  $command > /dev/null
}

# install keys
echo "adding keys..."
CreateAccount alice "slide moment original seven milk crawl help text kick fluid boring awkward doll wonder sure fragile plate grid hard next casual expire okay body"
CreateAccount bob "trap possible liquid elite embody host segment fantasy swim cable digital eager tiny broom burden diary earn hen grow engine pigeon fringe claim program"
CreateAccount carol "great immense still pill defense fetch pencil slow purchase symptom speed arm shoot fence have divorce cigar rapid hen vehicle pear evolve correct nerve"
CreateAccount dave "resist portion leisure lawn shell lunch inhale start pupil add fault segment hour dwarf toddler insect frame math remove twist easy paddle nice rhythm"
CreateAccount joe "diary behind allow lawn loop assault armor survey media toe rural mass skull way crater tobacco pair glare window soon lift pistol fitness annual"
CreateAccount steve "initial aisle blush gift shuffle hat solar dove dwarf soup credit acid update seed mercy thumb swear main verb image dove rose chat inmate"
CreateAccount olivia "dinosaur submit around jacket movie garden crash weather matter option time cabbage butter mind skin nature ride mosquito seat lock elephant merit empower come"

# commands
echo "creating coordinators..."
SpndCommand "profile create-coordinator" alice
SpndCommand "profile create-coordinator" bob
SpndCommand "profile create-coordinator" carol
SpndCommand "profile create-coordinator" dave
SpndCommand "profile create-coordinator" joe
SpndCommand "profile create-coordinator" steve
SpndCommand "profile create-coordinator" olivia

echo "creating chains..."
SpndCommand 'launch create-chain spn-10 https://github.com/tendermint/spn.git 0xaaa' alice
SpndCommand 'launch create-chain mars-1 https://github.com/lubtd/planet.git 0xbbb' bob
SpndCommand 'launch create-chain spn-11 https://github.com/tendermint/spn.git 0xccc' carol
SpndCommand 'launch create-chain spn-11 https://github.com/tendermint/spn.git 0xccc --account-balance 10000foo' carol

echo "add requests to chain 1"
SpndCommand 'launch request-add-account 1 1000foo,500bar' alice
SpndCommand 'launch request-add-account 1 10000foo,5500bar' bob
SpndCommand 'launch request-add-account 1 1000foo' carol
SpndCommand 'launch request-add-account 1 12938500bar' dave
SpndCommand 'launch request-add-account 1 1234baz' steve

echo "gentx content" > gentx.json
SpndCommand 'launch request-add-validator 1 gentx.json d/s5BWDjOgZAUT3Jp/MlPAS2Ob9w6mfnWiLYn0VelpE= 50foo aaa alice.com' alice
SpndCommand 'launch request-add-validator 1 gentx.json d/s5BWDjOgZAUT3Jp/MlPAS2Ob9w6mfnWiLYn0VelpE= 150foo bbb bob.com' bob
SpndCommand 'launch request-add-validator 1 gentx.json d/s5BWDjOgZAUT3Jp/MlPAS2Ob9w6mfnWiLYn0VelpE= 500foo ccc carol.com' carol
rm gentx.json

SpndCommand 'launch request-remove-validator 1 spn1aqn8ynvr3jmq67879qulzrwhchq5dtrvtx0nhe' dave
SpndCommand 'launch request-remove-account 1 spn1aqn8ynvr3jmq67879qulzrwhchq5dtrvtx0nhe' steve

echo "settle requests for chain 1"
SpndCommand 'launch settle-request approve 1 2' alice
SpndCommand 'launch settle-request approve 1 3' alice
SpndCommand 'launch settle-request reject 1 4' alice

SpndCommand 'launch settle-request approve 1 7' alice
SpndCommand 'launch settle-request reject 1 8' alice
