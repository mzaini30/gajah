function olah_wa(nomor_wa, teks) {
	teks = teks
		.split('\n')
		.map((x) => x.trimStart())
		.join('\n');
	teks = encodeURIComponent(teks);
	return `https://api.whatsapp.com/send?phone=${nomor_wa}&text=${teks}`;
}
