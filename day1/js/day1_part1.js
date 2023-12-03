import fs from 'fs';
import readline from 'readline';

var input_file = '../input.txt';
var r = readline.createInterface({
    input: fs.createReadStream(input_file)
});

let total = 0

r.on('line', function (text) {
    const lineWithoutLetters = text.replaceAll(/[a-z, A-Z]/g, '')
    if (!lineWithoutLetters.length) return
    else {
        total += parseInt(lineWithoutLetters[0] + lineWithoutLetters.slice(-1))
    }
}).on('close', () => {
    console.log(total);
})
