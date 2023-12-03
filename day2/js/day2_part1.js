import fs from 'fs';
import readline from 'readline';

const cubes = { red: 12, green: 13, blue: 14 }

export default function day2_part1() {
    var input_file = '../input.txt';
    var r = readline.createInterface({
        input: fs.createReadStream(input_file)
    });

    let total = 0

    r.on('line', function (text) {

        // Game number
        const gameNumber = text.split(':')[0].split(' ')[1]

        const gameSets = text.split(':')[1].split(';')

        let isGameValid = true
        gameSets.forEach(gameSet => {
            const colorValues = gameSet.split(',')
            colorValues.forEach(colorValue => {
                const color = colorValue.split(' ').slice(-1)
                const value = colorValue.split(' ')[1]
                
                if (cubes[color] < value) {
                    isGameValid = false
                }
            })
        })

        if(isGameValid) {
            total += parseInt(gameNumber)
        }

    }).on('close', () => {
        console.log("Part1: ", total);
    })
}
