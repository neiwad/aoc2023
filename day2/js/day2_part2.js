import fs from 'fs';
import readline from 'readline';

export default function day2_part2() {
    var input_file = '../input.txt';
    var r = readline.createInterface({
        input: fs.createReadStream(input_file)
    });

    let total = 0

    r.on('line', function (text) {

        // Game number
        const gameNumber = text.split(':')[0].split(' ')[1]

        const gameSets = text.split(':')[1].split(';')

        let gameSetMinColors = {
            red: 0, green: 0, blue: 0
        }
        gameSets.forEach(gameSet => {

            const colorValues = gameSet.split(',')
            colorValues.forEach(colorValue => {
                const color = colorValue.split(' ').slice(-1)[0]
                const value = colorValue.split(' ')[1]

                if (gameSetMinColors[color] < +value) {
                    gameSetMinColors[color] = +value
                }
            })



        })

        total += gameSetMinColors.red * gameSetMinColors.green * gameSetMinColors.blue

    }).on('close', () => {
        console.log("Part2: ", total);
    })
}
