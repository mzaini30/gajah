async function excalibur(server, data) {
	const body = new FormData();
	for (let n in data) {
		body.append(n, data[n]);
	}

	let proses = await fetch(
		server,
		{
			method: 'post',
			body
		},
		{
			'Content-Type': 'application/x-www-form-urlencoded'
		}
	);
	return proses;
}
