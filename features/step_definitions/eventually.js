module.exports = function (check, timeout = 1000) {
  const interval = 100
  const maxCount = timeout / interval

  return new Promise((resolve, reject) => {
    let count = 0;
    let checkerId = setInterval(() => {
      if (check()) {
        resolve()
        clearInterval(checkerId)
      } else {
        count += 1
        if (count >= maxCount) {
          reject("eventually did not resolve within the timeout period")
          clearInterval(checkerId)
        }
      }
    }, interval)
  })
}