
var games = ['1:0','2:0','3:0','4:0','2:1','3:1','4:1','3:2','4:2','4:3'];

/*
function points () {
    totalPoints = 0;
    
    for (var i = 0; i < games.length; i++) {
        var stringList = games[i].split(':');
        var a = stringList[0];
        var b = stringList[1];
        if (a > b) {
            totalPoints += 3;
        } else if (a == b) {
            totalPoints += 1;
        }
    }

    return totalPoints;
}
*/

function points (games) {
    let totalPoints = 0;
    games.map(game => {
        if (game[0] > game[2]) {
            totalPoints += 3;
        } else if (game[0] === game[2]) {
            totalPoints += 1;
        }
    });
    return totalPoints;
}

console.log(points(games));