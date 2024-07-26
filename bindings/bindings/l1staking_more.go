// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const L1StakingStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1026_storage\"},{\"astId\":1003,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1025_storage\"},{\"astId\":1005,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"_status\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1025_storage\"},{\"astId\":1007,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"rollupContract\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"stakingValue\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"withdrawalLockBlocks\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_uint256\"},{\"astId\":1010,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"rewardPercentage\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"gasLimitAddStaker\",\"offset\":0,\"slot\":\"155\",\"type\":\"t_uint256\"},{\"astId\":1012,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"gasLimitRemoveStakers\",\"offset\":0,\"slot\":\"156\",\"type\":\"t_uint256\"},{\"astId\":1013,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"slashRemaining\",\"offset\":0,\"slot\":\"157\",\"type\":\"t_uint256\"},{\"astId\":1014,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"whitelist\",\"offset\":0,\"slot\":\"158\",\"type\":\"t_mapping(t_address,t_bool)\"},{\"astId\":1015,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"removedList\",\"offset\":0,\"slot\":\"159\",\"type\":\"t_mapping(t_address,t_bool)\"},{\"astId\":1016,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"stakerSet\",\"offset\":0,\"slot\":\"160\",\"type\":\"t_array(t_address)1024_storage\"},{\"astId\":1017,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"stakerIndexes\",\"offset\":0,\"slot\":\"415\",\"type\":\"t_mapping(t_address,t_uint8)\"},{\"astId\":1018,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"deleteList\",\"offset\":0,\"slot\":\"416\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1019,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"deleteableHeight\",\"offset\":0,\"slot\":\"417\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1020,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"stakers\",\"offset\":0,\"slot\":\"418\",\"type\":\"t_mapping(t_address,t_struct(StakerInfo)1027_storage)\"},{\"astId\":1021,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"blsKeys\",\"offset\":0,\"slot\":\"419\",\"type\":\"t_mapping(t_bytes_memory_ptr,t_bool)\"},{\"astId\":1022,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"tmKeys\",\"offset\":0,\"slot\":\"420\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1023,\"contract\":\"contracts/l1/staking/L1Staking.sol:L1Staking\",\"label\":\"withdrawals\",\"offset\":0,\"slot\":\"421\",\"type\":\"t_mapping(t_address,t_uint256)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_address)1024_storage\":{\"encoding\":\"inplace\",\"label\":\"address[255]\",\"numberOfBytes\":\"8160\"},\"t_array(t_address)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"address[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1025_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1026_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_bytes_memory_ptr\":{\"encoding\":\"bytes\",\"label\":\"bytes\",\"numberOfBytes\":\"32\"},\"t_bytes_storage\":{\"encoding\":\"bytes\",\"label\":\"bytes\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"},\"t_mapping(t_address,t_struct(StakerInfo)1027_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct Types.StakerInfo)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(StakerInfo)1027_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_address,t_uint8)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint8)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint8\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes_memory_ptr,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes_memory_ptr\",\"value\":\"t_bool\"},\"t_struct(StakerInfo)1027_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.StakerInfo\",\"numberOfBytes\":\"96\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1StakingStorageLayout = new(solc.StorageLayout)

var L1StakingDeployedBin = "0x6080604052600436106102b6575f3560e01c80637a9262a211610170578063ab8c53dc116100d1578063d096c3c611610087578063df15503311610062578063df15503314610833578063e2a6ad5f14610877578063f2fde38b14610896575f80fd5b8063d096c3c6146107be578063d51c90a9146107dd578063dd4785f5146107f2575f80fd5b8063bfa02ba9116100b7578063bfa02ba914610754578063c7cd469a14610780578063cde4cd111461079f575f80fd5b8063ab8c53dc14610720578063ae81de5314610735575f80fd5b8063927ede2d116101265780639d48f4171161010c5780639d48f417146106b6578063a3066aab146106e2578063a4f209b014610701575f80fd5b8063927ede2d146106555780639b19251a14610688575f80fd5b80638a565ac3116101565780638a565ac3146105de5780638da5cb5b146105fd5780639168ae7214610627575f80fd5b80637a9262a21461057f578063831cfb58146105ab575f80fd5b806345bc4d101161021a5780636f1e8533116101d057806374fe27b7116101b657806374fe27b7146104ff578063797adbde146105255780637a4e87c314610544575f80fd5b80636f1e8533146104cc578063715018a6146104eb575f80fd5b80634d64903a116102005780634d64903a1461047357806352d472eb14610488578063692c565b1461049d575f80fd5b806345bc4d101461043557806345ff4c8014610454575f80fd5b80633cb747bf1161026f5780633ee2a1f9116102555780633ee2a1f9146103eb57806341de239b146103ff57806343352d6114610414575f80fd5b80633cb747bf146103a55780633ccfd60b146103d7575f80fd5b80632e407a6f1161029f5780632e407a6f146102f7578063303afb9e146103235780633a9bbede14610367575f80fd5b80632108db35146102ba5780632a28e5a3146102e2575b5f80fd5b3480156102c5575f80fd5b506102cf609c5481565b6040519081526020015b60405180910390f35b6102f56102f0366004613433565b6108b5565b005b348015610302575f80fd5b506102cf61031136600461349a565b6101a16020525f908152604090205481565b34801561032e575f80fd5b5061034261033d3660046134ba565b610d77565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102d9565b348015610372575f80fd5b5061039561038136600461349a565b609f6020525f908152604090205460ff1681565b60405190151581526020016102d9565b3480156103b0575f80fd5b507f0000000000000000000000000000000000000000000000000000000000000000610342565b3480156103e2575f80fd5b506102f5610da3565b3480156103f6575f80fd5b506102f561104c565b34801561040a575f80fd5b506102cf60995481565b34801561041f575f80fd5b50610428611372565b6040516102d991906134d1565b348015610440575f80fd5b506102cf61044f3660046134ba565b6113c5565b34801561045f575f80fd5b506102f561046e366004613518565b61178e565b34801561047e575f80fd5b506102cf60985481565b348015610493575f80fd5b506102cf609a5481565b3480156104a8575f80fd5b506103956104b73660046134ba565b6101a46020525f908152604090205460ff1681565b3480156104d7575f80fd5b506103956104e636600461349a565b611ce8565b3480156104f6575f80fd5b506102f5611d88565b34801561050a575f80fd5b506103956105193660046135a6565b60019695505050505050565b348015610530575f80fd5b5061034261053f3660046134ba565b611d9b565b34801561054f575f80fd5b5061039561055e36600461364c565b80516020818301810180516101a38252928201919093012091525460ff1681565b34801561058a575f80fd5b506102cf61059936600461349a565b6101a56020525f908152604090205481565b3480156105b6575f80fd5b506103427f000000000000000000000000000000000000000000000000000000000000000081565b3480156105e9575f80fd5b506102f56105f83660046134ba565b611dd1565b348015610608575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610342565b348015610632575f80fd5b5061064661064136600461349a565b611e96565b6040516102d9939291906136f1565b348015610660575f80fd5b506103427f000000000000000000000000000000000000000000000000000000000000000081565b348015610693575f80fd5b506103956106a236600461349a565b609e6020525f908152604090205460ff1681565b3480156106c1575f80fd5b506106d56106d03660046134ba565b611f56565b6040516102d9919061372e565b3480156106ed575f80fd5b506102f56106fc36600461349a565b612084565b34801561070c575f80fd5b506102f561071b3660046134ba565b612240565b34801561072b575f80fd5b506102cf609d5481565b348015610740575f80fd5b506102f561074f3660046134ba565b61230a565b34801561075f575f80fd5b506097546103429073ffffffffffffffffffffffffffffffffffffffff1681565b34801561078b575f80fd5b506102f561079a366004613787565b6123c7565b3480156107aa575f80fd5b506102f56107b936600461349a565b6125e5565b3480156107c9575f80fd5b506102cf6107d836600461349a565b612660565b3480156107e8575f80fd5b506102cf609b5481565b3480156107fd575f80fd5b5061082161080c36600461349a565b61019f6020525f908152604090205460ff1681565b60405160ff90911681526020016102d9565b34801561083e575f80fd5b5061039561084d36600461349a565b73ffffffffffffffffffffffffffffffffffffffff165f9081526101a16020526040902054151590565b348015610882575f80fd5b506102cf6108913660046137ee565b612701565b3480156108a1575f80fd5b506102f56108b036600461349a565b612875565b335f818152609e602052604090205460ff16610932576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f6e6f7420696e2077686974656c6973740000000000000000000000000000000060448201526064015b60405180910390fd5b335f9081526101a2602052604090205473ffffffffffffffffffffffffffffffffffffffff16156109bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f616c7265616479207265676973746572656400000000000000000000000000006044820152606401610929565b82158015906109dd57505f8381526101a4602052604090205460ff16155b610a43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f696e76616c69642074656e6465726d696e74207075626b6579000000000000006044820152606401610929565b8151610100148015610a7657506101a382604051610a61919061382d565b9081526040519081900360200190205460ff16155b610adc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c696420626c73207075626b657900000000000000000000000000006044820152606401610929565b6098543414610b47576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c6964207374616b696e672076616c756500000000000000000000006044820152606401610929565b6040518060600160405280610b593390565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083018790526040928301869052335f9081526101a28252839020845181547fffffffffffffffffffffffff00000000000000000000000000000000000000001693169290921782558301516001820155908201516002820190610bd990826138dd565b50905050610bec610be73390565b612929565b60016101a383604051610bff919061382d565b9081526040805191829003602090810190922080549315157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009485161790555f8681526101a49092529020805490911660011790557fb9c7babb56df9f2da4a30811a6c778e4e68af88b72712d56cf62c5516e20e199610c7c3390565b8484604051610c8d939291906136f1565b60405180910390a1335f9081526101a260209081526040918290208251606081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600281018054610d729484019190610ced90613848565b80601f0160208091040260200160405190810160405280929190818152602001828054610d1990613848565b8015610d645780601f10610d3b57610100808354040283529160200191610d64565b820191905f5260205f20905b815481529060010190602001808311610d4757829003601f168201915b505050505081525050612a96565b505050565b60a08160ff8110610d86575f80fd5b015473ffffffffffffffffffffffffffffffffffffffff16905081565b610dac33611ce8565b610e12576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6f6e6c79207374616b65720000000000000000000000000000000000000000006044820152606401610929565b335f9081526101a5602052604090205415610e89576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f7769746864726177696e670000000000000000000000000000000000000000006044820152606401610929565b610e923361084d565b15610ef9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f7374616b657220697320736c61736865640000000000000000000000000000006044820152606401610929565b609954610f069043613a26565b335f8181526101a56020526040902091909155610f2290612be9565b335f8181526101a560209081526040918290205491519182527f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5910160405180910390a26040805160018082528183019092525f916020808301908036833701905050905033815f81518110610f9a57610f9a613a39565b73ffffffffffffffffffffffffffffffffffffffff92909216602092830291909101820152335f908152609e8252604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00908116909155609f909352908190208054909216600117909155517f3f446646c03e618be8238a586960d6d625f35c653cdca1ef59609993e8ed2c849061103890839061372e565b60405180910390a161104981612d1c565b50565b611054612e3f565b5f5b6101a05481101561104957436101a15f6101a0848154811061107a5761107a613a39565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff1683528201929092526040019020541161136a5760a0600161019f5f6101a085815481106110cc576110cc613a39565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205461110a919060ff16613a66565b60ff1660ff811061111d5761111d613a39565b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556101a0805461019f915f918490811061115f5761115f613a39565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556101a080546101a1929190849081106111d0576111d0613a39565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff1683528201929092526040018120556101a0805461121290600190613a7f565b8154811061122257611222613a39565b5f918252602090912001546101a0805473ffffffffffffffffffffffffffffffffffffffff909216918390811061125b5761125b613a39565b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101a08054806112b2576112b2613a92565b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590556101a25f6101a083815481106112fc576112fc613a39565b5f91825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155600181018290559061136760028301826132f1565b50505b600101611056565b61137a613328565b60408051611fe08101918290529060a09060ff9082845b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611391575050505050905090565b6097545f9073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461145e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6f6e6c7920726f6c6c757020636f6e74726163740000000000000000000000006044820152606401610929565b611466612ec0565b5f61147083611f56565b90505f805b82518110156116a4575f6101a55f85848151811061149557611495613a39565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054111561154c576101a55f8483815181106114f3576114f3613a39565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9055609854826115459190613a26565b915061169c565b61159383828151811061156157611561613a39565b602002602001015173ffffffffffffffffffffffffffffffffffffffff165f9081526101a16020526040902054151590565b61169c576098546115a49083613a26565b91506115c88382815181106115bb576115bb613a39565b6020026020010151612be9565b609e5f8483815181106115dd576115dd613a39565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81549060ff02191690556001609f5f85848151811061164457611644613a39565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505b600101611475565b505f6064609a54836116b69190613abf565b6116c09190613ad6565b90506116cc8183613a7f565b609d5f8282546116dc9190613a26565b90915550506097546117049073ffffffffffffffffffffffffffffffffffffffff1682612f33565b7f654f4a61849f1b3ad10abb283d27f02d5fece7b820acc5a3b874713b58748b5a83604051611733919061372e565b60405180910390a17f3f446646c03e618be8238a586960d6d625f35c653cdca1ef59609993e8ed2c848360405161176a919061372e565b60405180910390a161177b83612d1c565b925050506117896001606555565b919050565b5f54610100900460ff16158080156117ac57505f54600160ff909116105b806117c55750303b1580156117c557505f5460ff166001145b611851576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610929565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156118ad575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff871661192a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c696420726f6c6c757020636f6e74726163740000000000000000006044820152606401610929565b5f8611611993576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c6964207374616b696e672076616c756500000000000000000000006044820152606401610929565b5f85116119fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f696e76616c6964207769746864726177616c206c6f636b20626c6f636b7300006044820152606401610929565b5f8311611a65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f696e76616c696420676173206c696d697420616464207374616b6572000000006044820152606401610929565b5f8211611ace576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f696e76616c696420676173206c696d69742072656d6f7665207374616b6572736044820152606401610929565b5f84118015611ade575060648411155b611b69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f696e76616c6964206368616c6c656e676572207265776172642070657263656e60448201527f74616765000000000000000000000000000000000000000000000000000000006064820152608401610929565b611b7161300a565b611b796130a8565b609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8916179055609a84905560988690556099859055609b839055609c829055604080515f8152602081018590527f0ef80fb82bf5439b8591744c0fde771f5d93dce7a7970b1e9b7828cdc3970e9c910160405180910390a1604080515f8152602081018490527fdd4b37d1f14888147fe4be2cfaedcbf148fd07ececc856d0433241b8b6e4a7f7910160405180910390a1604080515f8152602081018690527fa46de936426e045703b2d34a292a19fde92b329018db8e0da750033876b655ba910160405180910390a18015611cdf575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff81165f90815261019f602052604081205460ff168103611d1e57505f919050565b73ffffffffffffffffffffffffffffffffffffffff82165f81815261019f602052604090205460a090611d569060019060ff16613a66565b60ff1660ff8110611d6957611d69613a39565b015473ffffffffffffffffffffffffffffffffffffffff161492915050565b611d90612e3f565b611d995f613146565b565b6101a08181548110611dab575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b611dd9612e3f565b5f81118015611dea5750609c548114155b611e50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c6964206e657720676173206c696d697400000000000000000000006044820152606401610929565b609c80549082905560408051828152602081018490527fdd4b37d1f14888147fe4be2cfaedcbf148fd07ececc856d0433241b8b6e4a7f791015b60405180910390a15050565b6101a26020525f908152604090208054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff909316939192611ed590613848565b80601f0160208091040260200160405190810160405280929190818152602001828054611f0190613848565b8015611f4c5780601f10611f2357610100808354040283529160200191611f4c565b820191905f5260205f20905b815481529060010190602001808311611f2f57829003601f168201915b5050505050905083565b6060600182901c5f5b8115611f8657611f70816001613a26565b9050611f7d600183613a7f565b82169150611f5f565b8067ffffffffffffffff811115611f9f57611f9f61335f565b604051908082528060200260200182016040528015611fc8578160200160208202803683370190505b5092505f60015b60ff8160ff16101561207b57600160ff82161b8616156120735760a0611ff6600183613a66565b60ff1660ff811061200957612009613a39565b0154855173ffffffffffffffffffffffffffffffffffffffff9091169086908490811061203857612038613a39565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152612068826001613a26565b91508282101561207b575b600101611fcf565b50505050919050565b61208c612ec0565b335f9081526101a56020526040902054612102576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177616c206e6f742065786973740000000000000000000000006044820152606401610929565b335f9081526101a56020526040902054431161217a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f7769746864726177616c206c6f636b65640000000000000000000000000000006044820152606401610929565b335f9081526101a5602090815260408083208390556101a2909152812080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560018101829055906121d260028301826132f1565b50506121db3390565b60405173ffffffffffffffffffffffffffffffffffffffff838116825291909116907f89309c9b2aeaffbdce717113df9427298b20448c05919bf889e05f8c3094254b9060200160405180910390a261223681609854612f33565b6110496001606555565b612248612e3f565b5f81118015612258575060648111155b80156122665750609a548114155b6122cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f696e76616c6964207265776172642070657263656e74616765000000000000006044820152606401610929565b609a80549082905560408051828152602081018490527fa46de936426e045703b2d34a292a19fde92b329018db8e0da750033876b655ba9101611e8a565b612312612e3f565b5f811180156123235750609b548114155b612389576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c6964206e657720676173206c696d697400000000000000000000006044820152606401610929565b609b80549082905560408051828152602081018490527f0ef80fb82bf5439b8591744c0fde771f5d93dce7a7970b1e9b7828cdc3970e9c9101611e8a565b6123cf612e3f565b5f5b8381101561251957609f5f8686848181106123ee576123ee613a39565b9050602002016020810190612403919061349a565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f205460ff1615612493576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f696e2072656d6f766564206c69737400000000000000000000000000000000006044820152606401610929565b6001609e5f8787858181106124aa576124aa613a39565b90506020020160208101906124bf919061349a565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556001016123d1565b505f5b818110156125a157609e5f84848481811061253957612539613a39565b905060200201602081019061254e919061349a565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905560010161251c565b507fe375867e538b40218c1b3db2ccceaf875eb073e38b510449e6088c1539ac8622848484846040516125d79493929190613b60565b60405180910390a150505050565b6125ed612e3f565b6125f5612ec0565b609d546126028282612f33565b5f609d556040805173ffffffffffffffffffffffffffffffffffffffff84168152602081018390527fa1fefb6c5328a92a416e321ed50997303fe7135fd88c28b0592b21ce42b5cdd9910160405180910390a1506110496001606555565b5f61266a82611ce8565b6126d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e76616c6964207374616b65720000000000000000000000000000000000006044820152606401610929565b5073ffffffffffffffffffffffffffffffffffffffff165f90815261019f6020526040902054600160ff9091161b90565b5f60ff82111561276d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7374616b657273206c656e677468206f7574206f6620626f756e6473000000006044820152606401610929565b5f5b8281101561286d576127a184848381811061278c5761278c613a39565b90506020020160208101906104e6919061349a565b612807576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e76616c6964207374616b65720000000000000000000000000000000000006044820152606401610929565b61019f5f85858481811061281d5761281d613a39565b9050602002016020810190612832919061349a565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f2054600160ff90911681901b92909217910161276f565b505b92915050565b61287d612e3f565b73ffffffffffffffffffffffffffffffffffffffff8116612920576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610929565b61104981613146565b5f5b60ff8160ff161015612a33575f60a060ff80841690811061294e5761294e613a39565b015473ffffffffffffffffffffffffffffffffffffffff1603612a2b578160a08260ff1660ff811061298257612982613a39565b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556129d1816001613b91565b73ffffffffffffffffffffffffffffffffffffffff929092165f90815261019f6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff9093169290921790915550565b60010161292b565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f736c6f742066756c6c00000000000000000000000000000000000000000000006044820152606401610929565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b2267a7b7f00000000000000000000000000000000000000000000000000000000000000005f84604051602401612b059190613baa565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6d454d5100000000000000000000000000000000000000000000000000000000179052609b5490517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b168152612bb99493929190600401613bee565b5f604051808303815f87803b158015612bd0575f80fd5b505af1158015612be2573d5f803e3d5ffd5b5050505050565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526101a1602052604090205415612c76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f616c726561647920696e2064656c6574654c69737400000000000000000000006044820152606401610929565b6101a080546001810182555f919091527f7980fe0f714a613298681d64b7b8ffa7b148338dd52429f307d72798d5c317c40180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316179055609954612cf39043613a26565b73ffffffffffffffffffffffffffffffffffffffff9091165f9081526101a16020526040902055565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b2267a7b7f00000000000000000000000000000000000000000000000000000000000000005f84604051602401612d8b919061372e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0be67fcc00000000000000000000000000000000000000000000000000000000179052609c5490517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b168152612bb99493929190600401613bee565b60335473ffffffffffffffffffffffffffffffffffffffff163314611d99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610929565b600260655403612f2c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610929565b6002606555565b8015612fff575f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114612f8f576040519150601f19603f3d011682016040523d82523d5f602084013e612f94565b606091505b5050905080610d72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f526f6c6c75703a20455448207472616e73666572206661696c656400000000006044820152606401610929565b5050565b6001606555565b5f54610100900460ff166130a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610929565b611d996131bc565b5f54610100900460ff1661313e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610929565b611d9961325b565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16613252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610929565b611d9933613146565b5f54610100900460ff16613003576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610929565b5080546132fd90613848565b5f825580601f1061330c575050565b601f0160209004905f5260205f20908101906110499190613347565b60405180611fe0016040528060ff906020820280368337509192915050565b5b8082111561335b575f8155600101613348565b5090565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f83011261339b575f80fd5b813567ffffffffffffffff808211156133b6576133b661335f565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156133fc576133fc61335f565b81604052838152866020858801011115613414575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f8060408385031215613444575f80fd5b82359150602083013567ffffffffffffffff811115613461575f80fd5b61346d8582860161338c565b9150509250929050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611789575f80fd5b5f602082840312156134aa575f80fd5b6134b382613477565b9392505050565b5f602082840312156134ca575f80fd5b5035919050565b611fe0810181835f5b60ff81101561350f57815173ffffffffffffffffffffffffffffffffffffffff168352602092830192909101906001016134da565b50505092915050565b5f805f805f8060c0878903121561352d575f80fd5b61353687613477565b9860208801359850604088013597606081013597506080810135965060a00135945092505050565b5f8083601f84011261356e575f80fd5b50813567ffffffffffffffff811115613585575f80fd5b6020830191508360208260051b850101111561359f575f80fd5b9250929050565b5f805f805f80608087890312156135bb575f80fd5b86359550602087013567ffffffffffffffff808211156135d9575f80fd5b6135e58a838b0161355e565b9097509550604089013594506060890135915080821115613604575f80fd5b818901915089601f830112613617575f80fd5b813581811115613625575f80fd5b8a6020828501011115613636575f80fd5b6020830194508093505050509295509295509295565b5f6020828403121561365c575f80fd5b813567ffffffffffffffff811115613672575f80fd5b61367e8482850161338c565b949350505050565b5f5b838110156136a0578181015183820152602001613688565b50505f910152565b5f81518084526136bf816020860160208601613686565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f61372560608301846136a8565b95945050505050565b602080825282518282018190525f9190848201906040850190845b8181101561377b57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613749565b50909695505050505050565b5f805f806040858703121561379a575f80fd5b843567ffffffffffffffff808211156137b1575f80fd5b6137bd8883890161355e565b909650945060208701359150808211156137d5575f80fd5b506137e28782880161355e565b95989497509550505050565b5f80602083850312156137ff575f80fd5b823567ffffffffffffffff811115613815575f80fd5b6138218582860161355e565b90969095509350505050565b5f825161383e818460208701613686565b9190910192915050565b600181811c9082168061385c57607f821691505b602082108103613893577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b601f821115610d7257805f5260205f20601f840160051c810160208510156138be5750805b601f840160051c820191505b81811015612be2575f81556001016138ca565b815167ffffffffffffffff8111156138f7576138f761335f565b61390b816139058454613848565b84613899565b602080601f83116001811461395d575f84156139275750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556139f1565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156139a95788860151825594840194600190910190840161398a565b50858210156139e557878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561286f5761286f6139f9565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b60ff828116828216039081111561286f5761286f6139f9565b8181038181111561286f5761286f6139f9565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b808202811582820484141761286f5761286f6139f9565b5f82613b09577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b8183525f60208085019450825f5b85811015613b555773ffffffffffffffffffffffffffffffffffffffff613b4283613477565b1687529582019590820190600101613b1c565b509495945050505050565b604081525f613b73604083018688613b0e565b8281036020840152613b86818587613b0e565b979650505050505050565b60ff818116838216019081111561286f5761286f6139f9565b6020815273ffffffffffffffffffffffffffffffffffffffff8251166020820152602082015160408201525f604083015160608084015261367e60808401826136a8565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f613c2260808301856136a8565b90508260608301529594505050505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L1StakingStorageLayoutJSON), L1StakingStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1Staking"] = L1StakingStorageLayout
	deployedBytecodes["L1Staking"] = L1StakingDeployedBin
}
