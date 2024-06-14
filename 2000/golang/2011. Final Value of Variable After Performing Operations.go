func finalValueAfterOperations(operations []string) int {
    x:=0
    for _,opt:=range operations{
        if opt[1]=='+'{
            x++
        }else{
            x--
        }
    }
    return x
}
