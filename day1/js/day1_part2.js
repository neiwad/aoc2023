import fs from 'fs';
import readline from 'readline';

var input_file = '../input.txt';
var r = readline.createInterface({
    input: fs.createReadStream(input_file)
});

let total = 0
const numbers = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

r.on('line', function (text) {

    let newLine = ''
    text.split('').forEach((letter, index) => {
        // If letter is a number, add it to the new line
        if (parseInt(letter)) newLine += letter

        // If letter is a number spelled out, add the number to the new line
        numbers.forEach((number, i) => {
            const textToCheck = text.slice(index, index + number.length).toLowerCase()
            if (textToCheck === number) newLine += i + 1
        })
    })

    // Replace all other letters with nothing
    const lineWithoutLetters = newLine.replaceAll(/[a-z, A-Z]/g, '')

    // If line is empty, return
    if (lineWithoutLetters.length === 0) return
    else {
        total += parseInt(lineWithoutLetters[0] + lineWithoutLetters.slice(-1))
        return
    }
}).on('close', () => {
    console.log(total);
})
