package MaoNetworkEmulator

import MaoNetworkEmulatorAPI "MaoNetworkEmulator/MaoNetworkEmulator/api"

type Router struct {


}

func CreateRouter(deviceName string) *Router {
	// TODO
	return nil
}

func (r *Router) AddLink(interfaceName string, link BindLink) {
	// TODO
}

func (r *Router) AddLinkWithNumber(interfaceName string, interfaceNumber int, link BindLink) {
	// TODO
}



func (r *Router) AddProcessingModule(module *MaoNetworkEmulatorAPI.PacketProcessingModule) {
	// TODO
}

func (r *Router) SetAopChainBeforeFunc(hook MaoNetworkEmulatorAPI.AopChainBefore) {
	// TODO
}

func (r *Router) SetAopBeforeFunc(hook MaoNetworkEmulatorAPI.AopBefore) {
	// TODO
}

func (r *Router) SetAopAfterFunc(hook MaoNetworkEmulatorAPI.AopAfter) {
	// TODO
}

func (r *Router) SetAopChainAfterFunc(hook MaoNetworkEmulatorAPI.AopChainAfter) {
	// TODO
}