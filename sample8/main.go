package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type imageInfo struct {
	lineNo  int
	comment bool
	// image   string
	name string
	ver  string
}

// type LineInfo struct {
// 	lineNo int
// 	name   string
// }

func main() {
	fmt.Printf("")

	// var ss map[string][]string

	// fp, _ := os.Open("./compose.yml")
	// reader := bufio.NewReader(fp)

	// for {
	// 	line, err := reader.ReadString('\n')
	// 	fmt.Print(line)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		panic(err)
	// 	}
	// }
	repVer := map[string]string{
		"${test1}": "vaaa.00",
		"${test2}": "vbbb1.01",
	}

	imageLastLine := map[string]int{}
	ss := map[string][]imageInfo{}

	pattern := `(?P<comment>\s*#?\s*)image\s*:\s*(?P<name>[^\s]+)\s*:\s*(?P<ver>[^\s,]+)`
	// pattern := `image`
	re, _ := regexp.Compile(pattern)

	// lines := []string{
	// 	"    image: ${test1}:v1.00.01",
	// 	"#    image: ${test2}:v1.00.02 ",
	// 	"    #image: ${test3} :v1.00.03",
	// 	"  #  image: ${test3}: v1.00.04  # test",
	// }
	lines := []string{}
	fp, _ := os.Open("./compose.yml")
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		// lineはstring型で最後の改行が含まれない
		// fmt.Println(line)
	}

	for i, line := range lines {
		// match := re.FindStringSubmatch("    image: ${test1}:v1.00.01")
		match := re.FindStringSubmatch(line)
		if len(match) < 3 {
			continue
		}
		pp1 := match[re.SubexpIndex("comment")]
		pp2 := match[re.SubexpIndex("name")]
		pp3 := match[re.SubexpIndex("ver")]
		// fmt.Printf("%v, %v, %v\n", pp1, pp2, pp3)
		// _, exists := repVer[pp2]
		// if exists {
		comment := strings.Contains(pp1, "#")
		ss[pp2] = append(ss[pp2], imageInfo{
			lineNo:  i,
			comment: comment,
			// image:   pp1,
			name: pp2,
			ver:  pp3})
		imageLastLine[pp2] = i
		// }

		// val, exists := repVer[pp2]
		// if exists && val == pp3 {

		// 	lines[i] = fmt.Sprintf("    image:%v:%v", pp2, val)
		// } else {
		// 	lines[i] = fmt.Sprintf("    #image:%v:%v", pp2, val)
		// }

		// fmt.Printf("%v", match[0])
	}
	for k := range ss {
		ff := false
		imags := ss[k]
		for i := range imags {
			img := imags[i]
			_, exists := repVer[img.name]
			if exists {
				ff = true
				break
			}
		}
		if ff {
			for i := range imags {
				img := imags[i]
				index := img.lineNo

				v, exists := repVer[img.name]
				if exists && v == img.ver {
					lines[index] = fmt.Sprintf("    image: %v:%v", img.name, img.ver)
				} else {
					lines[index] = fmt.Sprintf("    #image: %v:%v", img.name, img.ver)
				}
			}
		}
	}

	str := strings.Join(lines, "\n")
	fmt.Printf("%v\n", str)
	// var ks []string
	// for k := range imageLastLine {
	// 	ks = append(ks, k)
	// }
	// slices.Reverse(ks)
	// fmt.Printf("%v\n", ss)

	f, _ := os.Create("./compose2.yml")
	defer f.Close()

	data := []byte(str)
	_, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}
	// fmt.Printf("%v\n", count)
}

// func existsItem(ss map[string][]imageInfo, name string) bool {
// 	values, exists := ss[name]
// 	for _, v := range values {
// 		v.comment
// 	}
// }

// func hasComment(ss map[string][]imageInfo, name string) bool {

// }
