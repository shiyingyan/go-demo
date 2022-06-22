package main



func main(){
	//fmt.Println("hello world")
	//fmt.Println(morestrings.ReverseRunes("Hello,Go!"))
	//fmt.Println(cmp.Diff("Hello World","Hello Go"))
	////fmt.Println(consts.DateTimeTemplate)
	//
	//a := []struct{in,out Count}{{"1","2"},{"3","4"}}
	//fmt.Println(a)
	//for index,v := range a {
	//	fmt.Println(index,v.in,v.out,v.in.ServeHTTP(string(v.out)))
	//}

}

type Handler interface {
    ServeHTTP(x string)
}

type Count string;

func (ctr Count) ServeHTTP(x string) string{
	return string(ctr) + x
}

