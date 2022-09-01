package MaoNetworkEmulatorAPI

type ModuleProcessResult uint

const (

	//MP_NORMAL Default. pass the packet to the next processing module
	MP_NORMAL ModuleProcessResult = iota + 1

	// MP_DROP force to drop the packet, stop processing by the processing chain.
	// AopAfter and AopChainAfter will be invoked.
	MP_DROP

	// MP_STOLEN the packet is consumed by the processing module. stop processing by the processing chain.
	// AopAfter and AopChainAfter will be invoked.
	MP_STOLEN
)

