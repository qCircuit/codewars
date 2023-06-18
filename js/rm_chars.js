/*
You are given a string containing a sequence of character sequences separated by commas.

Write a function which returns a new string containing the same character sequences except the first and the last ones but this time separated by spaces.

If the input string is empty or the removal of the first and last items would cause the resulting string to be empty, return an empty value (represented as a generic value NULL in the examples below).

"1,2,3"      =>  "2"
"1,2,3,4"    =>  "2 3"
"1,2,3,4,5"  =>  "2 3 4"

""     =>  NULL
"1"    =>  NULL
"1,2"  =>  NULL

*/

// mi solucion 
function array(string) {
    var lst = string.split(",")
    if (lst.length < 3) {
        return null
    }
    lst.shift()
    lst.pop()
    
    return lst.join(" ")
}

// el mejor
function array(arr){
    return arr.split(",").slice(1,-1).join(" ") || null;
  }