import bibit from './modul/olah-data.js'

import {driver} from 'selenia'
import fs  from 'fs'
import jadiLink from './modul/jadi-link.js'

import rapikan from './modul/rapikan.js'

let targetnya = fs.readFileSync('./config/target.csv', 'utf8').trim().split('\n')

async function app(){
	for (let x of bibit){
		for (let y of targetnya){
			await driver.get(y.replace(/{id}/, Math.random().toString().substr(2)))

			await driver.executeScript(`document.querySelector('[name="msg"]').value = '${jadiLink(x)}'`)
			await driver.executeScript(`document.querySelector('[name="write"]').click()`)

			const linknya = await driver.getCurrentUrl()

			const isiTeks = fs.readFileSync('./hasil.csv', 'utf8')
			fs.writeFileSync('./hasil.csv', `${isiTeks}\n${linknya}`)

			await driver.sleep(Math.random() * 2000)
		}
	}
	await rapikan()
}
app()
