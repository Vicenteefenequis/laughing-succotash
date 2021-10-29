export enum Time {
	Millisecond = 1,
	Second = 1000,
	Minute = 60000
}

export async function sleep(time: number) {
	return new Promise(resolve => setTimeout(resolve, time))
}