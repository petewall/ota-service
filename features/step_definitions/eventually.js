module.exports = function (check) {
  return new Promise((resolve, reject) => {
    let count = 0;
    let checkerId = setInterval(() => {
      if (check()) {
        resolve()
        clearInterval(checkerId)
      } else {
        count += 1
        if (count >= 10) {
          reject("eventually did not resolve within the timeout period")
          clearInterval(checkerId)
        }
      }
    }, 100)
  })
}