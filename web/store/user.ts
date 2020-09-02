import UserQuery from '@/apollo/queries/user.graphql'
import debug from 'debug'
import type { ActionTree, GetterTree, MutationTree } from 'vuex'
// https://typescript.nuxtjs.org/cookbook/store.html#vanilla

const d = debug('ml.store.user')

interface DistanceValue {
	value: number
	type: 'mile' | 'kilometer'
}
interface FluidValue {
	value: number
	type: 'gallon' | 'liter'
}
interface MoneyValue {
	value: number
	type: 'usd' | 'cad'
}

interface UserPreference {
	distance: DistanceValue['type'] | null
	fluid: FluidValue['type'] | null
	money: MoneyValue['type'] | null
	vehicle: Vehicle['id'] | null
}

interface Vehicle {
	id: string
	make: string
	model: string
	year: number
	odometer: DistanceValue
}

interface User {
	id: string
	email: string
	vehicles: Vehicle[] | null
	preference: UserPreference
}

export function state() {
	return {
		id: null as string | null,
		email: null as string | null,
		vehicles: null as Vehicle[] | null,
		preference: {} as UserPreference,
		selectedVehicle: null as Vehicle['id'] | null,
	}
}

export type RootState = ReturnType<typeof state>

export const mutations: MutationTree<RootState> = {
	setDefaultVehicle(state, vid: string) {
		if (state.vehicles?.find((v) => v.id === vid)) {
			state.selectedVehicle = vid
		}
	},
	setUser(state, { id, email, vehicles, preference }: User) {
		state.id = id
		state.email = email
		state.vehicles = vehicles
		state.preference = preference
		if (!state.selectedVehicle) state.selectedVehicle = preference.vehicle
	},
}

export const getters: GetterTree<RootState, RootState> = {
	selectedVehicle: (state) => {
		return state.vehicles?.find((v) => v.id === state.selectedVehicle)
	},
	hasVehicleSelected: (_, getters) => {
		return getters.selectedVehicle != null
	},
}

export const actions: ActionTree<RootState, RootState> = {
	async fetchUser({ commit }) {
		const client = this.app.apolloProvider.defaultClient
		let response = null
		try {
			response = await client.query({ query: UserQuery, variables: {} })
			commit('setUser', response.data.getUser as User)
		} catch (e) {
			d('there it goes, erroring again')
		}
	},
}