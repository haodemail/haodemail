let compare = function (pro, desc) {
  return function (obj1, obj2) {
    let val1 = obj1[pro];
    let val2 = obj2[pro];
    if (val1 < val2) {
      if (desc == false) {
        return -1;
      } else {
        return 1
      }
    } else if (val1 > val2) {
      if (desc == false) {
        return 1;
      } else {
        return -1
      }
    } else {
      return 0;
    }
  }
}
export {compare}
