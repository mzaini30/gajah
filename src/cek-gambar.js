export function cek(x){
	const pilih = x.match(/<img src=['"](.*?)\s*("(?:.*[^"])")?\s*['"]>/)
	let hasil = 'https://i.ibb.co/yYWd5t0/undraw-Waiting-for-you-ldha.png'
	if (pilih != null) {
		hasil = pilih[1]
	}
	return hasil
}