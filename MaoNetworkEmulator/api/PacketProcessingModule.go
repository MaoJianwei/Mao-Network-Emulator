package MaoNetworkEmulatorAPI

import "MaoNetworkEmulator/MaoNetworkEmulator"

type PacketProcessingModule interface {
	// ProcessPacket
	/**
	 *  router: router object as the context
	 *  packet: the packet to be processed
	 */
	ProcessPacket(router *MaoNetworkEmulator.Router, packet *MaoNetworkEmulator.Packet) ModuleProcessResult
}