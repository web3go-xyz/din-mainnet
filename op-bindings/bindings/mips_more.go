// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const MIPSStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/cannon/MIPS.sol:MIPS\",\"label\":\"oracle\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_contract(IOracle)1001\"}],\"types\":{\"t_contract(IOracle)1001\":{\"encoding\":\"inplace\",\"label\":\"contract IOracle\",\"numberOfBytes\":\"20\"}}}"

var MIPSStorageLayout = new(solc.StorageLayout)

var MIPSDeployedBin = "0x608060405234801561001057600080fd5b50600436106100415760003560e01c8063155633fe146100465780637dc0d1d01461006757806398bb138314610098575b600080fd5b61004e61016c565b6040805163ffffffff9092168252519081900360200190f35b61006f610174565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61015a600480360360408110156100ae57600080fd5b8101906020810181356401000000008111156100c957600080fd5b8201836020820111156100db57600080fd5b803590602001918460018302840111640100000000831117156100fd57600080fd5b91939092909160208101903564010000000081111561011b57600080fd5b82018360208201111561012d57600080fd5b8035906020019184600183028401116401000000008311171561014f57600080fd5b509092509050610190565b60408051918252519081900360200190f35b634000000081565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600061019a611a62565b608081146101a757600080fd5b604051610600146101b757600080fd5b606486146101c457600080fd5b61016684146101d257600080fd5b6101ef565b8035602084810360031b9190911c8352920192910190565b8560806101fe602082846101d7565b9150915061020e602082846101d7565b9150915061021e600482846101d7565b9150915061022e600482846101d7565b9150915061023e600482846101d7565b9150915061024e600482846101d7565b9150915061025e600482846101d7565b9150915061026e600482846101d7565b9150915061027e600182846101d7565b9150915061028e600182846101d7565b9150915061029e600882846101d7565b6020810190819052909250905060005b60208110156102d0576102c3600483856101d7565b90935091506001016102ae565b505050806101200151156102ee576102e6610710565b915050610708565b6101408101805160010167ffffffffffffffff1690526060810151600090610316908261081e565b9050603f601a82901c16600281148061033557508063ffffffff166003145b15610382576103788163ffffffff1660021461035257601f610355565b60005b60ff16600261036b856303ffffff16601a6108e6565b63ffffffff16901b610959565b9350505050610708565b6101608301516000908190601f601086901c81169190601587901c16602081106103a857fe5b602002015192508063ffffffff851615806103c957508463ffffffff16601c145b156103fa578661016001518263ffffffff16602081106103e557fe5b6020020151925050601f600b86901c166104b1565b60208563ffffffff16101561045d578463ffffffff16600c148061042457508463ffffffff16600d145b8061043557508463ffffffff16600e145b15610446578561ffff169250610458565b6104558661ffff1660106108e6565b92505b6104b1565b60288563ffffffff1610158061047957508463ffffffff166022145b8061048a57508463ffffffff166026145b156104b1578661016001518263ffffffff16602081106104a657fe5b602002015192508190505b60048563ffffffff16101580156104ce575060088563ffffffff16105b806104df57508463ffffffff166001145b156104fe576104f0858784876109c4565b975050505050505050610708565b63ffffffff60006020878316106105635761051e8861ffff1660106108e6565b9095019463fffffffc861661053481600161081e565b915060288863ffffffff161015801561055457508763ffffffff16603014155b1561056157809250600093505b505b600061057189888885610b4d565b63ffffffff9081169150603f8a16908916158015610596575060088163ffffffff1610155b80156105a85750601c8163ffffffff16105b15610687578063ffffffff16600814806105c857508063ffffffff166009145b156105ff576105ed8163ffffffff166008146105e457856105e7565b60005b89610959565b9b505050505050505050505050610708565b8063ffffffff16600a1415610620576105ed858963ffffffff8a1615611213565b8063ffffffff16600b1415610642576105ed858963ffffffff8a161515611213565b8063ffffffff16600c1415610659576105ed6112f8565b60108163ffffffff16101580156106765750601c8163ffffffff16105b15610687576105ed81898988611770565b8863ffffffff1660381480156106a2575063ffffffff861615155b156106d15760018b61016001518763ffffffff16602081106106c057fe5b63ffffffff90921660209290920201525b8363ffffffff1663ffffffff146106ee576106ee84600184611954565b6106fa85836001611213565b9b5050505050505050505050505b949350505050565b6000610728565b602083810382015183520192910190565b60806040518061073a60208285610717565b9150925061074a60208285610717565b9150925061075a60048285610717565b9150925061076a60048285610717565b9150925061077a60048285610717565b9150925061078a60048285610717565b9150925061079a60048285610717565b915092506107aa60048285610717565b915092506107ba60018285610717565b915092506107ca60018285610717565b915092506107da60088285610717565b60209091019350905060005b6020811015610808576107fb60048386610717565b90945091506001016107e6565b506000815281810382a081900390209150505b90565b60008061082a836119f0565b9050600384161561083a57600080fd5b602081019035610857565b60009081526020919091526040902090565b8460051c8160005b601b8110156108af5760208501943583821c60011680156108875760018114610898576108a5565b6108918285610845565b93506108a5565b6108a28483610845565b93505b505060010161085f565b5060805191508181146108ca57630badf00d60005260206000fd5b5050601f8516601c0360031b1c63ffffffff1691505092915050565b600063ffffffff8381167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80850183169190911c821615159160016020869003821681901b830191861691821b92911b0182610943576000610945565b815b90861663ffffffff16179250505092915050565b6000610963611a62565b5060e08051610100805163ffffffff90811690935284831690526080918516156109b357806008018261016001518663ffffffff16602081106109a257fe5b63ffffffff90921660209290920201525b6109bb610710565b95945050505050565b60006109ce611a62565b5060806000600463ffffffff881614806109ee57508663ffffffff166005145b15610a645760008261016001518663ffffffff1660208110610a0c57fe5b602002015190508063ffffffff168563ffffffff16148015610a3457508763ffffffff166004145b80610a5c57508063ffffffff168563ffffffff1614158015610a5c57508763ffffffff166005145b915050610ae1565b8663ffffffff1660061415610a825760008460030b13159050610ae1565b8663ffffffff1660071415610a9f5760008460030b139050610ae1565b8663ffffffff1660011415610ae157601f601087901c1680610ac55760008560030b1291505b8063ffffffff1660011415610adf5760008560030b121591505b505b606082018051608084015163ffffffff169091528115610b27576002610b0c8861ffff1660106108e6565b63ffffffff90811690911b8201600401166080840152610b39565b60808301805160040163ffffffff1690525b610b41610710565b98975050505050505050565b6000603f601a86901c81169086166020821015610f215760088263ffffffff1610158015610b815750600f8263ffffffff16105b15610c28578163ffffffff1660081415610b9d57506020610c23565b8163ffffffff1660091415610bb457506021610c23565b8163ffffffff16600a1415610bcb5750602a610c23565b8163ffffffff16600b1415610be25750602b610c23565b8163ffffffff16600c1415610bf957506024610c23565b8163ffffffff16600d1415610c1057506025610c23565b8163ffffffff16600e1415610c23575060265b600091505b63ffffffff8216610e7157601f600688901c16602063ffffffff83161015610d455760088263ffffffff1610610c6357869350505050610708565b63ffffffff8216610c835763ffffffff86811691161b9250610708915050565b8163ffffffff1660021415610ca75763ffffffff86811691161c9250610708915050565b8163ffffffff1660031415610cd2576103788163ffffffff168763ffffffff16901c826020036108e6565b8163ffffffff1660041415610cf6575050505063ffffffff8216601f84161b610708565b8163ffffffff1660061415610d1a575050505063ffffffff8216601f84161c610708565b8163ffffffff1660071415610d45576103788763ffffffff168763ffffffff16901c886020036108e6565b8163ffffffff1660201480610d6057508163ffffffff166021145b15610d72578587019350505050610708565b8163ffffffff1660221480610d8d57508163ffffffff166023145b15610d9f578587039350505050610708565b8163ffffffff1660241415610dbb578587169350505050610708565b8163ffffffff1660251415610dd7578587179350505050610708565b8163ffffffff1660261415610df3578587189350505050610708565b8163ffffffff1660271415610e0f575050505082821719610708565b8163ffffffff16602a1415610e42578560030b8760030b12610e32576000610e35565b60015b60ff169350505050610708565b8163ffffffff16602b1415610e6b578563ffffffff168763ffffffff1610610e32576000610e35565b50610f1c565b8163ffffffff16600f1415610e945760108563ffffffff16901b92505050610708565b8163ffffffff16601c1415610f1c578063ffffffff1660021415610ebd57505050828202610708565b8063ffffffff1660201480610ed857508063ffffffff166021145b15610f1c578063ffffffff1660201415610ef0579419945b60005b6380000000871615610f12576401fffffffe600197881b169601610ef3565b9250610708915050565b6111ac565b60288263ffffffff16101561108b578163ffffffff1660201415610f6e57610f658660031660080260180363ffffffff168563ffffffff16901c60ff1660086108e6565b92505050610708565b8163ffffffff1660211415610fa457610f658660021660080260100363ffffffff168563ffffffff16901c61ffff1660106108e6565b8163ffffffff1660221415610fd55750505063ffffffff60086003851602811681811b198416918316901b17610708565b8163ffffffff1660231415610fee578392505050610708565b8163ffffffff1660241415611022578560031660080260180363ffffffff168463ffffffff16901c60ff1692505050610708565b8163ffffffff1660251415611057578560021660080260100363ffffffff168463ffffffff16901c61ffff1692505050610708565b8163ffffffff1660261415610f1c5750505063ffffffff60086003851602601803811681811c198416918316901c17610708565b8163ffffffff16602814156110c35750505060ff63ffffffff60086003861602601803811682811b9091188316918416901b17610708565b8163ffffffff16602914156110fc5750505061ffff63ffffffff60086002861602601003811682811b9091188316918416901b17610708565b8163ffffffff16602a141561112d5750505063ffffffff60086003851602811681811c198316918416901c17610708565b8163ffffffff16602b1415611146578492505050610708565b8163ffffffff16602e141561117a5750505063ffffffff60086003851602601803811681811b198316918416901b17610708565b8163ffffffff1660301415611193578392505050610708565b8163ffffffff16603814156111ac578492505050610708565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f696e76616c696420696e737472756374696f6e00000000000000000000000000604482015290519081900360640190fd5b600061121d611a62565b506080602063ffffffff86161061129557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f76616c6964207265676973746572000000000000000000000000000000000000604482015290519081900360640190fd5b63ffffffff8516158015906112a75750825b156112d557838161016001518663ffffffff16602081106112c457fe5b63ffffffff90921660209290920201525b60808101805163ffffffff808216606085015260049091011690526109bb610710565b6000611302611a62565b506101e051604081015160808083015160a084015160c09094015191936000928392919063ffffffff8616610ffa141561137a5781610fff81161561134c57610fff811661100003015b63ffffffff84166113705760e08801805163ffffffff838201169091529550611374565b8395505b50611723565b8563ffffffff16610fcd14156113965763400000009450611723565b8563ffffffff1661101814156113af5760019450611723565b8563ffffffff1661109614156113e757600161012088015260ff83166101008801526113d9610710565b97505050505050505061081b565b8563ffffffff16610fa314156115a15763ffffffff83166114075761159c565b63ffffffff8316600514156115795760006114298363fffffffc16600161081e565b6000805460208b01516040808d015181517fe03110e1000000000000000000000000000000000000000000000000000000008152600481019390935263ffffffff16602483015280519495509293849373ffffffffffffffffffffffffffffffffffffffff9093169263e03110e19260448082019391829003018186803b1580156114b357600080fd5b505afa1580156114c7573d6000803e3d6000fd5b505050506040513d60408110156114dd57600080fd5b508051602090910151909250905060038516600481900382811015611500578092505b508185101561150d578491505b8260088302610100031c925082600882021b9250600180600883600403021b036001806008858560040301021b0391508119811690508381198616179450505061155f8563fffffffc16600185611954565b60408a018051820163ffffffff169052965061159c915050565b63ffffffff8316600314156115905780945061159c565b63ffffffff9450600993505b611723565b8563ffffffff16610fa414156116755763ffffffff8316600114806115cc575063ffffffff83166002145b806115dd575063ffffffff83166004145b156115ea5780945061159c565b63ffffffff83166006141561159057600061160c8363fffffffc16600161081e565b60208901519091506003841660040383811015611627578093505b83900360089081029290921c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600193850293841b0116911b1760208801526000604088015293508361159c565b8563ffffffff16610fd71415611723578163ffffffff16600314156117175763ffffffff831615806116ad575063ffffffff83166005145b806116be575063ffffffff83166003145b156116cc576000945061159c565b63ffffffff8316600114806116e7575063ffffffff83166002145b806116f8575063ffffffff83166006145b80611709575063ffffffff83166004145b15611590576001945061159c565b63ffffffff9450601693505b6101608701805163ffffffff808816604090920191909152905185821660e09091015260808801805180831660608b01526004019091169052611764610710565b97505050505050505090565b600061177a611a62565b5060806000601063ffffffff88161415611799575060c08101516118f1565b8663ffffffff16601114156117b95763ffffffff861660c08301526118f1565b8663ffffffff16601214156117d3575060a08101516118f1565b8663ffffffff16601314156117f35763ffffffff861660a08301526118f1565b8663ffffffff16601814156118285763ffffffff600387810b9087900b02602081901c821660c08501521660a08301526118f1565b8663ffffffff166019141561185a5763ffffffff86811681871602602081901c821660c08501521660a08301526118f1565b8663ffffffff16601a14156118a5578460030b8660030b8161187857fe5b0763ffffffff1660c0830152600385810b9087900b8161189457fe5b0563ffffffff1660a08301526118f1565b8663ffffffff16601b14156118f1578463ffffffff168663ffffffff16816118c957fe5b0663ffffffff90811660c0840152858116908716816118e457fe5b0463ffffffff1660a08301525b63ffffffff84161561192657808261016001518563ffffffff166020811061191557fe5b63ffffffff90921660209290920201525b60808201805163ffffffff80821660608601526004909101169052611949610710565b979650505050505050565b600061195f836119f0565b9050600384161561196f57600080fd5b6020810190601f8516601c0360031b83811b913563ffffffff90911b1916178460051c60005b601b8110156119e55760208401933582821c60011680156119bd57600181146119ce576119db565b6119c78286610845565b94506119db565b6119d88583610845565b94505b5050600101611995565b505060805250505050565b60ff81166103800261016681019036906104e601811015611a5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611aed6023913960400191505060405180910390fd5b50919050565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101919091526101608101611ac8611acd565b905290565b604051806104000160405280602090602082028036833750919291505056fe636865636b207468617420746865726520697320656e6f7567682063616c6c64617461a164736f6c6343000706000a"

var MIPSDeployedSourceMap = "1017:22386:0:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1415:45;;;:::i;:::-;;;;;;;;;;;;;;;;;;;1783:21;;;:::i;:::-;;;;;;;;;;;;;;;;;;;14339:4789;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;14339:4789:0;;-1:-1:-1;14339:4789:0;-1:-1:-1;14339:4789:0;:::i;:::-;;;;;;;;;;;;;;;;1415:45;1450:10;1415:45;:::o;1783:21::-;;;;;;:::o;14339:4789::-;14417:7;14432:18;;:::i;:::-;14532:4;14525:5;14522:15;14512:2;;14593:1;14591;14584:11;14512:2;14629:4;14623:11;14636;14620:28;14610:2;;14694:1;14692;14685:11;14610:2;14742:3;14724:16;14721:25;14711:2;;14808:1;14806;14799:11;14711:2;14852:3;14838:12;14835:21;14825:2;;14917:1;14915;14908:11;14825:2;14934:370;;;15152:24;;15140:2;15136:13;;;15133:1;15129:21;15125:52;;;;15186:20;;15232:21;;;15278:18;;;15012:292::o;:::-;15320:16;15371:4;15404:18;15419:2;15416:1;15413;15404:18;:::i;:::-;15396:26;;;;15448:18;15463:2;15460:1;15457;15448:18;:::i;:::-;15440:26;;;;15496:17;15511:1;15508;15505;15496:17;:::i;:::-;15488:25;;;;15546:17;15561:1;15558;15555;15546:17;:::i;:::-;15538:25;;;;15584:17;15599:1;15596;15593;15584:17;:::i;:::-;15576:25;;;;15626:17;15641:1;15638;15635;15626:17;:::i;:::-;15618:25;;;;15664:17;15679:1;15676;15673;15664:17;:::i;:::-;15656:25;;;;15702:17;15717:1;15714;15711;15702:17;:::i;:::-;15694:25;;;;15742:17;15757:1;15754;15751;15742:17;:::i;:::-;15734:25;;;;15786:17;15801:1;15798;15795;15786:17;:::i;:::-;15778:25;;;;15828:17;15843:1;15840;15837;15828:17;:::i;:::-;15877:2;15870:10;;15860:21;;;;15820:25;;-1:-1:-1;15870:10:0;-1:-1:-1;15948:1:0;15933:77;15958:2;15955:1;15952:9;15933:77;;;15991:17;16006:1;16003;16000;15991:17;:::i;:::-;15983:25;;-1:-1:-1;15983:25:0;-1:-1:-1;15976:1:0;15969:9;15933:77;;;15937:14;;;16037:5;:12;;;16034:86;;;16100:13;:11;:13::i;:::-;16093:20;;;;;16034:86;16125:10;;;:15;;16139:1;16125:15;;;;;16194:8;;;;-1:-1:-1;;16186:20:0;;-1:-1:-1;16186:7:0;:20::i;:::-;16172:34;-1:-1:-1;16229:10:0;16237:2;16229:10;;;;16290:1;16280:11;;;:26;;;16295:6;:11;;16305:1;16295:11;16280:26;16276:308;;;16515:62;16526:6;:11;;16536:1;16526:11;:20;;16544:2;16526:20;;;16540:1;16526:20;16515:62;;16575:1;16548:23;16551:4;16556:10;16551:15;16568:2;16548;:23::i;:::-;:28;;;;16515:10;:62::i;:::-;16508:69;;;;;;;16276:308;16785:15;;;;16612:9;;;;16733:4;16727:2;16719:10;;;16718:19;;;16785:15;16810:2;16802:10;;;16801:19;16785:36;;;;;;;;;;;;-1:-1:-1;16842:5:0;16857:11;;;;;:29;;;16872:6;:14;;16882:4;16872:14;16857:29;16853:636;;;16929:5;:15;;;16945:5;16929:22;;;;;;;;;;;;;;-1:-1:-1;;16982:4:0;16976:2;16968:10;;;16967:19;16853:636;;;17012:4;17003:6;:13;;;16999:490;;;17103:6;:13;;17113:3;17103:13;:30;;;;17120:6;:13;;17130:3;17120:13;17103:30;:47;;;;17137:6;:13;;17147:3;17137:13;17103:47;17099:181;;;17189:4;17194:6;17189:11;17184:16;;17099:181;;;17252:19;17255:4;17260:6;17255:11;17268:2;17252;:19::i;:::-;17247:24;;17099:181;16999:490;;;17306:4;17296:6;:14;;;;:32;;;;17314:6;:14;;17324:4;17314:14;17296:32;:50;;;;17332:6;:14;;17342:4;17332:14;17296:50;17292:197;;;17396:5;:15;;;17412:5;17396:22;;;;;;;;;;;;;17391:27;;17477:5;17469:13;;17292:197;17510:1;17500:6;:11;;;;:25;;;;;17524:1;17515:6;:10;;;17500:25;17499:42;;;;17530:6;:11;;17540:1;17530:11;17499:42;17495:107;;;17558:37;17571:6;17579:4;17585:5;17592:2;17558:12;:37::i;:::-;17551:44;;;;;;;;;;;17495:107;17627:13;17608:16;17747:4;17737:14;;;;17733:328;;17796:19;17799:4;17804:6;17799:11;17812:2;17796;:19::i;:::-;17790:25;;;;17842:10;17837:15;;17866:16;17837:15;17880:1;17866:7;:16::i;:::-;17860:22;;17904:4;17894:6;:14;;;;:32;;;;;17912:6;:14;;17922:4;17912:14;;17894:32;17890:165;;;17967:4;17955:16;;18045:1;18037:9;;17890:165;17733:328;;18078:10;18091:26;18099:4;18105:2;18109;18113:3;18091:7;:26::i;:::-;18120:10;18091:39;;;;-1:-1:-1;18208:4:0;18201:11;;;18232;;;:24;;;;;18255:1;18247:4;:9;;;;18232:24;:39;;;;;18267:4;18260;:11;;;18232:39;18228:589;;;18285:4;:9;;18293:1;18285:9;:22;;;;18298:4;:9;;18306:1;18298:9;18285:22;18281:102;;;18337:37;18348:4;:9;;18356:1;18348:9;:21;;18364:5;18348:21;;;18360:1;18348:21;18371:2;18337:10;:37::i;:::-;18330:44;;;;;;;;;;;;;;;18281:102;18395:4;:11;;18403:3;18395:11;18391:79;;;18433:28;18442:5;18449:2;18453:7;;;;18433:8;:28::i;18391:79::-;18481:4;:11;;18489:3;18481:11;18477:79;;;18519:28;18528:5;18535:2;18539:7;;;;;18519:8;:28::i;18477:79::-;18606:4;:11;;18614:3;18606:11;18602:58;;;18636:15;:13;:15::i;18602:58::-;18733:4;18725;:12;;;;:27;;;;;18748:4;18741;:11;;;18725:27;18721:90;;;18771:31;18782:4;18788:2;18792;18796:5;18771:10;:31::i;18721:90::-;18861:6;:14;;18871:4;18861:14;:28;;;;-1:-1:-1;18879:10:0;;;;;18861:28;18857:75;;;18924:1;18899:5;:15;;;18915:5;18899:22;;;;;;;;;:26;;;;:22;;;;;;:26;18857:75;18962:9;:26;;18975:13;18962:26;18958:74;;18998:27;19007:9;19018:1;19021:3;18998:8;:27::i;:::-;19097:26;19106:5;19113:3;19118:4;19097:8;:26::i;:::-;19090:33;;;;;;;;;;;;;14339:4789;;;;;;;:::o;2069:1331::-;2110:11;2238:176;;;2330:2;2326:13;;;2316:24;;2310:31;2299:43;;2362:13;;2393;;;2289:125::o;:::-;2433:4;2472;2466:11;2510:5;2534:21;2552:2;2548;2542:4;2534:21;:::i;:::-;2522:33;;;;2585:21;2603:2;2599;2593:4;2585:21;:::i;:::-;2573:33;;;;2640:20;2658:1;2654:2;2648:4;2640:20;:::i;:::-;2628:32;;;;2697:20;2715:1;2711:2;2705:4;2697:20;:::i;:::-;2685:32;;;;2742:20;2760:1;2756:2;2750:4;2742:20;:::i;:::-;2730:32;;;;2791:20;2809:1;2805:2;2799:4;2791:20;:::i;:::-;2779:32;;;;2836:20;2854:1;2850:2;2844:4;2836:20;:::i;:::-;2824:32;;;;2881:20;2899:1;2895:2;2889:4;2881:20;:::i;:::-;2869:32;;;;2928:20;2946:1;2942:2;2936:4;2928:20;:::i;:::-;2916:32;;;;2979:20;2997:1;2993:2;2987:4;2979:20;:::i;:::-;2967:32;;;;3028:20;3046:1;3042:2;3036:4;3028:20;:::i;:::-;3081:2;3071:13;;;;-1:-1:-1;3016:32:0;-1:-1:-1;3129:1:0;3114:84;3139:2;3136:1;3133:9;3114:84;;;3176:20;3194:1;3190:2;3184:4;3176:20;:::i;:::-;3164:32;;-1:-1:-1;3164:32:0;-1:-1:-1;3157:1:0;3150:9;3114:84;;;3118:14;3229:1;3225:2;3218:13;3274:5;3270:2;3266:14;3259:5;3254:27;3359:14;;;3342:32;;;-1:-1:-1;;2069:1331:0;;:::o;11741:1270::-;11812:10;11830:14;11847:23;11859:10;11847:11;:23::i;:::-;11830:40;;11906:1;11900:4;11896:12;11893:2;;;11921:1;11918;11911:12;11893:2;12023;12011:15;;;11974:20;12033:169;;;;12072:12;;;12158:2;12151:13;;;;12191:2;12178:16;;;12062:140::o;:::-;12228:4;12225:1;12221:12;12252:4;12375:1;12360:273;12385:2;12382:1;12379:9;12360:273;;;12484:2;12472:15;;;12433:20;12507:12;;;12521:1;12503:20;12532:42;;;;12588:1;12583:42;;;;12496:129;;12532:42;12549:23;12564:7;12558:4;12549:23;:::i;:::-;12541:31;;12532:42;;12583;12600:23;12618:4;12609:7;12600:23;:::i;:::-;12592:31;;12496:129;-1:-1:-1;;12403:1:0;12396:9;12360:273;;;12364:14;12661:4;12655:11;12640:26;;12730:7;12724:4;12721:17;12711:2;;12787:10;12784:1;12777:21;12817:2;12814:1;12807:13;12711:2;-1:-1:-1;;12933:2:0;12923:13;;12911:10;12907:30;12904:1;12900:38;12956:16;12974:10;12952:33;;-1:-1:-1;;11741:1270:0;;;;:::o;1809:256::-;1868:6;1899:14;;;;1907:5;;;;1899:14;;;;;;1898:21;;;;;1911:1;1950:2;:6;;;1944:13;;;;;1943:19;;1942:28;;;;;;;1992:8;;1991:14;1898:21;2037;;2057:1;2037:21;;;2048:6;2037:21;2025:8;;;;;:34;;-1:-1:-1;;;1809:256:0;;;;:::o;10413:401::-;10480:7;10495:18;;:::i;:::-;-1:-1:-1;10576:8:0;;;10601:12;;;10590:23;;;;;;;10619:19;;;;;10545:4;;10648:12;;;10644:140;;10697:6;10704:1;10697:8;10670:5;:15;;;10686:7;10670:24;;;;;;;;;:35;;;;:24;;;;;;:35;10644:140;10796:13;:11;:13::i;:::-;10789:20;10413:401;-1:-1:-1;;;;;10413:401:0:o;8255:1063::-;8348:7;8363:18;;:::i;:::-;-1:-1:-1;8413:4:0;8428:17;8474:1;8464:11;;;;;:26;;;8479:6;:11;;8489:1;8479:11;8464:26;8460:514;;;8513:9;8525:5;:15;;;8541:5;8525:22;;;;;;;;;;;;;8513:34;;8577:2;8571:8;;:2;:8;;;:23;;;;;8583:6;:11;;8593:1;8583:11;8571:23;8570:54;;;;8606:2;8600:8;;:2;:8;;;;:23;;;;;8612:6;:11;;8622:1;8612:11;8600:23;8555:69;;8460:514;;;;8641:6;:11;;8651:1;8641:11;8637:337;;;8684:1;8677:2;8671:14;;;;8656:29;;8637:337;;;8710:6;:11;;8720:1;8710:11;8706:268;;;8752:1;8746:2;8740:13;;;8725:28;;8706:268;;;8778:6;:11;;8788:1;8778:11;8774:200;;;8844:4;8838:2;8830:10;;;8829:19;8861:8;8857:42;;8898:1;8892:2;8886:13;;;8871:28;;8857:42;8920:3;:8;;8927:1;8920:8;8916:43;;;8958:1;8951:2;8945:14;;;;8930:29;;8916:43;8774:200;;8996:8;;;;;9021:12;;;;9010:23;;;;;9071:216;;;;9147:1;9126:19;9129:4;9134:6;9129:11;9142:2;9126;:19::i;:::-;:22;;;;;;;9112:37;;9121:1;9112:37;9097:52;:12;;;:52;9071:216;;;9244:12;;;;;9259:1;9244:16;9229:31;;;;9071:216;9300:13;:11;:13::i;:::-;9293:20;8255:1063;-1:-1:-1;;;;;;;;8255:1063:0:o;19132:4269::-;19219:6;19249:10;19257:2;19249:10;;;;;;19292:11;;19388:4;19379:13;;19375:3986;;;19489:1;19479:6;:11;;;;:27;;;;;19503:3;19494:6;:12;;;19479:27;19475:462;;;19522:6;:11;;19532:1;19522:11;19518:383;;;-1:-1:-1;19544:4:0;19518:383;;;19584:6;:11;;19594:1;19584:11;19580:321;;;-1:-1:-1;19606:4:0;19580:321;;;19642:6;:13;;19652:3;19642:13;19638:263;;;-1:-1:-1;19666:4:0;19638:263;;;19699:6;:13;;19709:3;19699:13;19695:206;;;-1:-1:-1;19723:4:0;19695:206;;;19757:6;:13;;19767:3;19757:13;19753:148;;;-1:-1:-1;19781:4:0;19753:148;;;19814:6;:13;;19824:3;19814:13;19810:91;;;-1:-1:-1;19838:4:0;19810:91;;;19870:6;:13;;19880:3;19870:13;19866:35;;;-1:-1:-1;19894:4:0;19866:35;19927:1;19918:10;;19475:462;19978:11;;;19974:1701;;20030:4;20025:1;20017:9;;;20016:18;20055:4;20017:9;20048:11;;;20044:588;;;20085:4;20077;:12;;;20073:549;;20100:2;20093:9;;;;;;;20073:549;20180:12;;;20176:446;;20203:11;;;;;;;;-1:-1:-1;20196:18:0;;-1:-1:-1;;20196:18:0;20176:446;20249:4;:12;;20257:4;20249:12;20245:377;;;20272:11;;;;;;;;-1:-1:-1;20265:18:0;;-1:-1:-1;;20265:18:0;20245:377;20318:4;:12;;20326:4;20318:12;20314:308;;;20341:25;20350:5;20344:11;;:2;:11;;;;20360:5;20357:2;:8;20341:2;:25::i;20314:308::-;20401:4;:12;;20409:4;20401:12;20397:225;;;-1:-1:-1;;;;20424:15:0;;;20434:4;20431:7;;20424:15;20417:22;;20397:225;20478:4;:12;;20486:4;20478:12;20474:148;;;-1:-1:-1;;;;20501:15:0;;;20511:4;20508:7;;20501:15;20494:22;;20474:148;20555:4;:12;;20563:4;20555:12;20551:71;;;20578:19;20587:2;20581:8;;:2;:8;;;;20594:2;20591;:5;20578:2;:19::i;20551:71::-;20720:4;:12;;20728:4;20720:12;:28;;;;20736:4;:12;;20744:4;20736:12;20720:28;20716:602;;;20762:2;20759;:5;20752:12;;;;;;;20716:602;20802:4;:12;;20810:4;20802:12;:28;;;;20818:4;:12;;20826:4;20818:12;20802:28;20798:520;;;20844:2;20841;:5;20834:12;;;;;;;20798:520;20884:4;:12;;20892:4;20884:12;20880:438;;;20910:2;20907;:5;20900:12;;;;;;;20880:438;20951:4;:12;;20959:4;20951:12;20947:371;;;20978:2;20975;:5;20967:14;;;;;;;20947:371;21017:4;:12;;21025:4;21017:12;21013:305;;;21044:2;21041;:5;21033:14;;;;;;;21013:305;21084:4;:12;;21092:4;21084:12;21080:238;;;-1:-1:-1;;;;21109:5:0;;;21107:8;21100:15;;21080:238;21151:4;:12;;21159:4;21151:12;21147:171;;;21200:2;21184:19;;21190:2;21184:19;;;:27;;21210:1;21184:27;;;21206:1;21184:27;21177:34;;;;;;;;;21147:171;21239:4;:12;;21247:4;21239:12;21235:83;;;21275:2;21272:5;;:2;:5;;;:13;;21284:1;21272:13;;21235:83;19974:1701;;;;21336:6;:13;;21346:3;21336:13;21332:343;;;21364:2;21360;:6;;;;21353:13;;;;;;21332:343;21392:6;:14;;21402:4;21392:14;21388:287;;;21435:4;:9;;21443:1;21435:9;21431:49;;;-1:-1:-1;;;21460:19:0;;;21446:34;;21431:49;21501:4;:12;;21509:4;21501:12;:28;;;;21517:4;:12;;21525:4;21517:12;21501:28;21497:170;;;21554:4;:12;;21562:4;21554:12;21550:26;;;21573:3;;;21550:26;21588:8;21602:45;21612:10;21609:13;;:18;21602:45;;21636:8;21631:3;21636:8;;;;;21631:3;21602:45;;;21655:1;-1:-1:-1;21648:8:0;;-1:-1:-1;;21648:8:0;21497:170;19375:3986;;;21700:4;21691:6;:13;;;21687:1674;;;21718:6;:14;;21728:4;21718:14;21714:776;;;21758:36;21774:2;21777:1;21774:4;21780:1;21773:8;21770:2;:11;21762:20;;:3;:20;;;;21786:4;21761:29;21792:1;21758:2;:36::i;:::-;21751:43;;;;;;21714:776;21813:6;:14;;21823:4;21813:14;21809:681;;;21853:39;21869:2;21872:1;21869:4;21875:1;21868:8;21865:2;:11;21857:20;;:3;:20;;;;21881:6;21856:31;21889:2;21853;:39::i;21809:681::-;21911:6;:14;;21921:4;21911:14;21907:583;;;-1:-1:-1;;;21958:17:0;21973:1;21970;21967:4;;21966:8;21958:17;;21999:32;;;22054:5;22049:10;;21958:17;;;;;22048:18;22041:25;;21907:583;22085:6;:14;;22095:4;22085:14;22081:409;;;22110:3;22103:10;;;;;;22081:409;22140:6;:14;;22150:4;22140:14;22136:354;;;22194:2;22197:1;22194:4;22200:1;22193:8;22190:2;:11;22182:20;;:3;:20;;;;22206:4;22181:29;22174:36;;;;;;22136:354;22229:6;:14;;22239:4;22229:14;22225:265;;;22283:2;22286:1;22283:4;22289:1;22282:8;22279:2;:11;22271:20;;:3;:20;;;;22295:6;22270:31;22263:38;;;;;;22225:265;22320:6;:14;;22330:4;22320:14;22316:174;;;-1:-1:-1;;;22367:20:0;22385:1;22382;22379:4;;22378:8;22375:2;:11;22367:20;;22411:35;;;22469:5;22464:10;;22367:20;;;;;22463:18;22456:25;;21687:1674;22506:6;:14;;22516:4;22506:14;22502:859;;;-1:-1:-1;;;22553:4:0;22549:26;22573:1;22570;22567:4;;22566:8;22563:2;:11;22549:26;;22617:21;;;22597:42;;;22655:10;;22550:7;;;22549:26;;22654:18;22647:25;;22502:859;22689:6;:14;;22699:4;22689:14;22685:676;;;-1:-1:-1;;;22736:6:0;22732:28;22758:1;22755;22752:4;;22751:8;22748:2;:11;22732:28;;22802:23;;;22782:44;;;22842:10;;22733:9;;;22732:28;;22841:18;22834:25;;22685:676;22876:6;:14;;22886:4;22876:14;22872:489;;;-1:-1:-1;;;22921:16:0;22935:1;22932;22929:4;;22928:8;22921:16;;22959:32;;;23013:5;23007:11;;22921:16;;;;;23006:19;22999:26;;22872:489;23042:6;:14;;23052:4;23042:14;23038:323;;;23079:2;23072:9;;;;;;23038:323;23098:6;:14;;23108:4;23098:14;23094:267;;;-1:-1:-1;;;23143:19:0;23160:1;23157;23154:4;;23153:8;23150:2;:11;23143:19;;23184:35;;;23241:5;23235:11;;23143:19;;;;;23234;23227:26;;23094:267;23270:6;:14;;23280:4;23270:14;23266:95;;;23295:3;23288:10;;;;;;23266:95;23321:6;:14;;23331:4;23321:14;23317:44;;;23346:2;23339:9;;;;;;23317:44;23367:29;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;10818:455;10901:7;10916:18;;:::i;:::-;-1:-1:-1;10966:4:0;11000:2;10989:13;;;;10981:40;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;11099:13;;;;;;;:28;;;11116:11;11099:28;11095:80;;;11165:3;11137:5;:15;;;11153:8;11137:25;;;;;;;;;:31;;;;:25;;;;;;:31;11095:80;11192:12;;;;;11181:23;;;;:8;;;:23;11240:1;11225:16;;;11210:31;;;11255:13;:11;:13::i;3404:4847::-;3447:7;3462:18;;:::i;:::-;-1:-1:-1;3547:15:0;;:18;;;;3512:4;3622:18;;;;3658;;;;3694;;;;;3512:4;;3527:17;;;;3622:18;3658;3723;;;3737:4;3723:18;3719:4375;;;3777:2;3794:4;3791:7;;:12;3787:98;;3871:4;3868:7;;3860:4;:16;3854:22;3787:98;3896:7;;;3892:105;;3920:10;;;;;3940:16;;;;;;;;3920:10;-1:-1:-1;3892:105:0;;;3986:2;3981:7;;3892:105;3719:4375;;;;4013:10;:18;;4027:4;4013:18;4009:4085;;;1450:10;4054:14;;4009:4085;;;4085:10;:18;;4099:4;4085:18;4081:4013;;;4149:1;4144:6;;4081:4013;;;4167:10;:18;;4181:4;4167:18;4163:3931;;;4230:4;4215:12;;;:19;4242:26;;;:14;;;:26;4283:13;:11;:13::i;:::-;4276:20;;;;;;;;;;;4163:3931;4313:10;:18;;4327:4;4313:18;4309:3785;;;4442:14;;;4438:1768;;;;;4533:22;;;1671:1;4533:22;4529:1677;;;4654:10;4667:27;4675:2;4680:10;4675:15;4692:1;4667:7;:27::i;:::-;4745:11;4776:6;;4796:17;;;;4815:20;;;;;4776:60;;;;;;;;;;;;;;;;;;;;4654:40;;-1:-1:-1;4745:11:0;;;;4776:6;;;;;:19;;:60;;;;;;;;;;;:6;:60;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;-1:-1:-1;4776:60:0;;;;;;;;;-1:-1:-1;4776:60:0;-1:-1:-1;4949:1:0;4941:10;;5029:1;5025:17;;;5090;;;5087:2;;;5120:5;5110:15;;5087:2;;5189:6;5185:2;5182:14;5179:2;;;5209;5199:12;;5179:2;5301:3;5296:1;5288:6;5284:14;5279:3;5275:24;5271:34;5264:41;;5366:3;5362:1;5351:9;5347:17;5343:27;5336:34;;5476:1;5472;5468;5456:9;5453:1;5449:17;5445:25;5441:33;5437:41;5589:1;5585;5581;5572:6;5560:9;5557:1;5553:17;5549:30;5545:38;5541:46;5537:54;5519:72;;5675:10;5671:15;5665:4;5661:26;5653:34;;5777:3;5769:4;5765:9;5760:3;5756:19;5753:28;5746:35;;;;5857:33;5866:2;5871:10;5866:15;5883:1;5886:3;5857:8;:33::i;:::-;5900:20;;;:38;;;;;;;;;-1:-1:-1;4529:1677:0;;-1:-1:-1;;4529:1677:0;;5986:18;;;1594:1;5986:18;5982:224;;;6135:2;6130:7;;5982:224;;;6167:10;6162:15;;1742:3;6187:10;;5982:224;4309:3785;;;6222:10;:18;;6236:4;6222:18;6218:1876;;;6355:15;;;1525:1;6355:15;;:34;;-1:-1:-1;6374:15:0;;;1558:1;6374:15;6355:34;:57;;;-1:-1:-1;6393:19:0;;;1631:1;6393:19;6355:57;6351:1172;;;6429:2;6424:7;;6351:1172;;;6493:23;;;1712:1;6493:23;6489:1034;;;6548:10;6561:27;6569:2;6574:10;6569:15;6586:1;6561:7;:27::i;:::-;6652:17;;;;6548:40;;-1:-1:-1;6782:1:0;6774:10;;6862:1;6858:17;6923:13;;;6920:2;;;6945:5;6939:11;;6920:2;7189:14;;;7023:1;7185:22;;;7181:32;;;;7092:26;7116:1;7015:10;;;7096:18;;;7092:26;7177:43;7011:20;;7271:12;7321:17;;;:23;7377:1;7354:20;;;:24;7019:2;-1:-1:-1;7019:2:0;6489:1034;;6218:1876;7539:10;:18;;7553:4;7539:18;7535:559;;;7613:2;:7;;7619:1;7613:7;7609:479;;;7674:14;;;;;:40;;-1:-1:-1;7692:22:0;;;1671:1;7692:22;7674:40;:62;;;-1:-1:-1;7718:18:0;;;1594:1;7718:18;7674:62;7670:312;;;7755:1;7750:6;;7670:312;;;7789:15;;;1525:1;7789:15;;:34;;-1:-1:-1;7808:15:0;;;1558:1;7808:15;7789:34;:61;;;-1:-1:-1;7827:23:0;;;1712:1;7827:23;7789:61;:84;;;-1:-1:-1;7854:19:0;;;1631:1;7854:19;7789:84;7785:197;;;7892:1;7887:6;;7785:197;;7609:479;8011:10;8006:15;;1774:4;8031:11;;7609:479;8100:15;;;;;:23;;;;:18;;;;:23;;;;8129:15;;:23;;;:18;;;;:23;-1:-1:-1;8170:12:0;;;;8159:23;;;:8;;;:23;8218:1;8203:16;8188:31;;;;;8233:13;:11;:13::i;:::-;8226:20;;;;;;;;;3404:4847;:::o;9322:1087::-;9412:7;9427:18;;:::i;:::-;-1:-1:-1;9477:4:0;9492:10;9520:4;9512:12;;;;9508:732;;;-1:-1:-1;9532:8:0;;;;9508:732;;;9563:4;:12;;9571:4;9563:12;9559:681;;;9577:13;;;:8;;;:13;9559:681;;;9613:4;:12;;9621:4;9613:12;9609:631;;;-1:-1:-1;9633:8:0;;;;9609:631;;;9664:4;:12;;9672:4;9664:12;9660:580;;;9678:13;;;:8;;;:13;9660:580;;;9714:4;:12;;9722:4;9714:12;9710:530;;;9824:7;9781:16;9764;;;9781;;;;9764:33;9829:2;9824:7;;;;;9806:8;;;:26;9840:22;:8;;;:22;9710:530;;;9879:4;:12;;9887:4;9879:12;9875:365;;;9941:10;9930;;;9941;;;9930:21;9983:2;9978:7;;;;;9960:8;;;:26;9994:22;:8;;;:22;9875:365;;;10033:4;:12;;10041:4;10033:12;10029:211;;;10096:2;10080:19;;10086:2;10080:19;;;;;;;;10062:38;;:8;;;:38;10126:19;;;;;;;;;;;;;;10108:38;;:8;;;:38;10029:211;;;10163:4;:12;;10171:4;10163:12;10159:81;;;10207:2;10204:5;;:2;:5;;;;;;;;10193:16;;;;:8;;;:16;10228:5;;;;;;;;;;;;10217:16;;:8;;;:16;10159:81;10250:13;;;;10246:65;;10301:3;10273:5;:15;;;10289:8;10273:25;;;;;;;;;:31;;;;:25;;;;;;:31;10246:65;10328:12;;;;;10317:23;;;;:8;;;:23;10376:1;10361:16;;;10346:31;;;10391:13;:11;:13::i;:::-;10384:20;9322:1087;-1:-1:-1;;;;;;;9322:1087:0:o;13134:1145::-;13217:14;13234:23;13246:10;13234:11;:23::i;:::-;13217:40;;13293:1;13287:4;13283:12;13280:2;;;13308:1;13305;13298:12;13280:2;13416;13579:15;;;13434:2;13424:13;;13412:10;13408:30;13405:1;13401:38;13544:17;;;13361:20;;13529:10;13518:22;;;13514:27;13504:38;13501:61;13796:4;13793:1;13789:12;13943:1;13928:273;13953:2;13950:1;13947:9;13928:273;;;14052:2;14040:15;;;14001:20;14075:12;;;14089:1;14071:20;14100:42;;;;14156:1;14151:42;;;;14064:129;;14100:42;14117:23;14132:7;14126:4;14117:23;:::i;:::-;14109:31;;14100:42;;14151;14168:23;14186:4;14177:7;14168:23;:::i;:::-;14160:31;;14064:129;-1:-1:-1;;13971:1:0;13964:9;13928:273;;;-1:-1:-1;;14215:4:0;14208:18;-1:-1:-1;;;;13272:1003:0:o;11277:460::-;11552:19;;;11575:5;11552:29;11545:3;:37;;;11623:14;;11658;;11652:21;;;11644:69;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;11719:13;11277:460;;;:::o;-1:-1:-1:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;;;;:::o;:::-;;;;;;;;;;;;;;;;;;;;;;;;:::o"

func init() {
	if err := json.Unmarshal([]byte(MIPSStorageLayoutJSON), MIPSStorageLayout); err != nil {
		panic(err)
	}

	layouts["MIPS"] = MIPSStorageLayout
	deployedBytecodes["MIPS"] = MIPSDeployedBin
}
