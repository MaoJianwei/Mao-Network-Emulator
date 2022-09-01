package MaoNetworkEmulatorAPI

import "MaoNetworkEmulator/MaoNetworkEmulator"

// AopChainBefore
/**
 *  router: router object as the context
 *  packet: the packet being processed
 */
type AopChainBefore func(router *MaoNetworkEmulator.Router, packet *MaoNetworkEmulator.Packet)

// AopBefore
/**
 *  processModuleInfo: name of the processing function, and the package of the processing module
 *  router: router object as the context
 *  packet: the packet being processed
 */
type AopBefore func(processModuleInfo string, router *MaoNetworkEmulator.Router, packet *MaoNetworkEmulator.Packet)

// AopAfter
/**
 *  processModuleInfo: name of the processing function, and the package of the processing module
 *  processResult: result judged by this processing module
 *  router: router object as the context
 *  packet: the packet being processed
 */
type AopAfter func(processModuleInfo string, processResult ModuleProcessResult, router *MaoNetworkEmulator.Router, packet *MaoNetworkEmulator.Packet)

// AopChainAfter
/**
 *  lastProcessModuleInfo: name of the processing function, and the package of the last one processing module
 *  processResult: result judged by this last one processing module
 *  router: router object as the context
 *  packet: the packet being processed
 */
type AopChainAfter func(lastProcessModuleInfo string, lastProcessResult ModuleProcessResult, router *MaoNetworkEmulator.Router, packet *MaoNetworkEmulator.Packet)



