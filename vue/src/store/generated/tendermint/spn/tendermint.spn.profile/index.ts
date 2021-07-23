import { txClient, queryClient, MissingWalletError } from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'

import { Coordinator } from "./module/types/profile/coordinator"
import { CoordinatorDescription } from "./module/types/profile/coordinator"
import { CoordinatorByAddress } from "./module/types/profile/coordinator"
import { QueryAllCoordinatorByAddressRequest } from "./module/types/profile/query"
import { QueryAllCoordinatorByAddressResponse } from "./module/types/profile/query"
import { ValidatorByAddress } from "./module/types/profile/validator"
import { ValidatorDescription } from "./module/types/profile/validator"
import { ValidatorByConsAddress } from "./module/types/profile/validator"
import { ConsensusKeyNonce } from "./module/types/profile/validator"


export { Coordinator, CoordinatorDescription, CoordinatorByAddress, QueryAllCoordinatorByAddressRequest, QueryAllCoordinatorByAddressResponse, ValidatorByAddress, ValidatorDescription, ValidatorByConsAddress, ConsensusKeyNonce };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				ConsensusKeyNonce: {},
				ValidatorByConsAddress: {},
				ValidatorByAddress: {},
				ValidatorByAddressAll: {},
				CoordinatorByAddress: {},
				Coordinator: {},
				CoordinatorAll: {},
				
				_Structure: {
						Coordinator: getStructure(Coordinator.fromPartial({})),
						CoordinatorDescription: getStructure(CoordinatorDescription.fromPartial({})),
						CoordinatorByAddress: getStructure(CoordinatorByAddress.fromPartial({})),
						QueryAllCoordinatorByAddressRequest: getStructure(QueryAllCoordinatorByAddressRequest.fromPartial({})),
						QueryAllCoordinatorByAddressResponse: getStructure(QueryAllCoordinatorByAddressResponse.fromPartial({})),
						ValidatorByAddress: getStructure(ValidatorByAddress.fromPartial({})),
						ValidatorDescription: getStructure(ValidatorDescription.fromPartial({})),
						ValidatorByConsAddress: getStructure(ValidatorByConsAddress.fromPartial({})),
						ConsensusKeyNonce: getStructure(ConsensusKeyNonce.fromPartial({})),
						
		},
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(subscription)
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(subscription)
		}
	},
	getters: {
				getConsensusKeyNonce: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ConsensusKeyNonce[JSON.stringify(params)] ?? {}
		},
				getValidatorByConsAddress: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorByConsAddress[JSON.stringify(params)] ?? {}
		},
				getValidatorByAddress: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorByAddress[JSON.stringify(params)] ?? {}
		},
				getValidatorByAddressAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorByAddressAll[JSON.stringify(params)] ?? {}
		},
				getCoordinatorByAddress: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CoordinatorByAddress[JSON.stringify(params)] ?? {}
		},
				getCoordinator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Coordinator[JSON.stringify(params)] ?? {}
		},
				getCoordinatorAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CoordinatorAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: tendermint.spn.profile initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					await dispatch(subscription.action, subscription.payload)
				}catch(e) {
					throw new SpVuexError('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryConsensusKeyNonce({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryConsensusKeyNonce( key.consAddress)).data
				
					
				commit('QUERY', { query: 'ConsensusKeyNonce', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryConsensusKeyNonce', payload: { options: { all }, params: {...key},query }})
				return getters['getConsensusKeyNonce']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryConsensusKeyNonce', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorByConsAddress({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryValidatorByConsAddress( key.consAddress)).data
				
					
				commit('QUERY', { query: 'ValidatorByConsAddress', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorByConsAddress', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorByConsAddress']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryValidatorByConsAddress', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorByAddress({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryValidatorByAddress( key.address)).data
				
					
				commit('QUERY', { query: 'ValidatorByAddress', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorByAddress', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorByAddress']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryValidatorByAddress', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorByAddressAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryValidatorByAddressAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.nextKey!=null) {
					let next_values=(await queryClient.queryValidatorByAddressAll({...query, 'pagination.key':(<any> value).pagination.nextKey})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValidatorByAddressAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorByAddressAll', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorByAddressAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryValidatorByAddressAll', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCoordinatorByAddress({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCoordinatorByAddress( key.address)).data
				
					
				commit('QUERY', { query: 'CoordinatorByAddress', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCoordinatorByAddress', payload: { options: { all }, params: {...key},query }})
				return getters['getCoordinatorByAddress']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryCoordinatorByAddress', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCoordinator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCoordinator( key.id)).data
				
					
				commit('QUERY', { query: 'Coordinator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCoordinator', payload: { options: { all }, params: {...key},query }})
				return getters['getCoordinator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryCoordinator', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCoordinatorAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params: {...key}, query=null }) {
			try {
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCoordinatorAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.nextKey!=null) {
					let next_values=(await queryClient.queryCoordinatorAll({...query, 'pagination.key':(<any> value).pagination.nextKey})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CoordinatorAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCoordinatorAll', payload: { options: { all }, params: {...key},query }})
				return getters['getCoordinatorAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryCoordinatorAll', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateCoordinator({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateCoordinator(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateCoordinator:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateCoordinator:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateCoordinator({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateCoordinator(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateCoordinator:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateCoordinator:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		
	}
}
