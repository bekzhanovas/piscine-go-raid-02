    package main

    import (
    	"fmt"
    	//"strconv"
    	//"cast/ToInt"
    )

    func main() {
    	str1 := ".96.4...1"
    	str2 := "1...6...4"
    	str3 := "5.481.39."
    	str4 := "..795..43"
    	str5 := ".3..8...."
    	str6 := "4.5.23.18"
    	str7 := ".1.63..59"
    	str8 := ".59.7.83."
    	str9 := "..359...7"
    	raw_data := [9]string{str1, str2, str3, str4, str5, str6, str7, str8, str9}
    	//переводим точки на 0 и в int
    	for i, numbers := range raw_data {
    		raw_data[i] = dots_to_zero(numbers)
    	}
    	////переводим string в int
    	int_data := str_to_int(raw_data)

    	//fmt.Println(int_data)
    	//fmt.Println(int_data[1][0])

    	for row := range int_data {
    		fmt.Println(int_data[row])
    	}

    	/*a3 := check_square(int_data, 0, 3)
    	a1 := check_row(int_data, 0)
    	a2 := check_column(int_data, 3)
    	fmt.Println(a1)
    	fmt.Println(a2)
    	fmt.Println(a3)
    	d1 := drop_repeated_values(a1, a2)
    	fmt.Println(d1)
    	d2 := drop_repeated_values(d1, a3)
    	fmt.Println(d2)*/

    	for {
    		counter := 0
    		for row := 0; row < 9; row++ {
    			for column := 0; column < 9; column++ {
    				if int_data[row][column] == 0 {
    					d1 := drop_repeated_values(drop_repeated_values(check_row(int_data, row), check_column(int_data, column)), check_square(int_data, row, column))
    					//fmt.Println(d1)
    					if len(d1) == 1 {
    						int_data[row][column] = d1[0]
    						counter++
    					}
    				}
    			}
    		}
    		if counter == 0 {
    			break
    		}
    	}

    	fmt.Println()
    	for row := range int_data {
    		fmt.Println(int_data[row])
    	}
    	fmt.Println()
    	
    	for row := 0; row < 9; row++ {
    		for column := 0; column < 9; column++ {
    			if int_data[row][column] == 0 {
    				d1 := drop_repeated_values(drop_repeated_values(check_row(int_data, row), check_column(int_data, column)), check_square(int_data, row, column))
    				fmt.Println(d1)
    				if len(d1) == 1 {
    					int_data[row][column] = d1[0]

    				}
    			}
    		}
    	}

    }

    func drop_repeated_values(slice1 []int, slice2 []int) []int {
    	//slice1 := []int{8, 2, 3, 7, 5}
    	//slice2 := []int{9, 2, 3, 7, 8, 6}
    	//fmt.Println(slice1, slice2)
    	slice := make([]int, 9)
    	for i, number1 := range slice1 {
    		counter := 0
    		for _, number2 := range slice2 {
    			if number1 == number2 {
    				counter++
    			}
    		}
    		//fmt.Println(counter, slice1)
    		if counter > 0 {

    			slice[i] = slice1[i]
    			/*slice1[i] = slice1[len(slice1)-1] // Copy last element to index i.
    			slice1[len(slice1)-1] = 0         // Erase last element (write zero value).
    			slice1 = slice1[:len(slice1)-1]   // Truncate slice.
    			*/
    		}

    	}
    	counter := 0
    	for i := 0; i < 9; i++ {
    		if slice[i] != 0 {
    			counter++
    		}
    	}
    	slice_fin := make([]int, counter)
    	//fmt.Println(slice)
    	i := 0
    	for j := 0; j < 9; j++ {

    		if slice[j] != 0 {
    			slice_fin[i] = slice[j]
    			i++
    		}
    	}

    	//fmt.Println(slice_fin )
    	return slice_fin
    }

    func check_square(int_data [][]int, row int, column int) []int {
    	all_variables := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    	//fmt.Println(all_variables)
    	//int_data
    	//i
    	for i := (row / 3) * 3; i < (row/3)*3+3; i++ {
    		for k := (column / 3) * 3; k < (column/3)*3+3; k++ {
    			data := int_data[i][k]
    			for j := 0; j < len(all_variables); j++ {
    				if data == all_variables[j] {
    					all_variables[j] = all_variables[len(all_variables)-1] // Copy last element to index i.
    					all_variables[len(all_variables)-1] = 0                // Erase last element (write zero value).
    					all_variables = all_variables[:len(all_variables)-1]   // Truncate slice.
    				}
    			}
    		}
    	}
    	//fmt.Println(all_variables)
    	return all_variables
    }

    func check_row(int_data [][]int, row int) []int {
    	all_variables := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    	//fmt.Println(all_variables)
    	//int_data
    	//i
    	for i := 0; i < 9; i++ {
    		data := int_data[row][i]
    		for j := 0; j < len(all_variables); j++ {
    			if data == all_variables[j] {
    				all_variables[j] = all_variables[len(all_variables)-1] // Copy last element to index i.
    				all_variables[len(all_variables)-1] = 0                // Erase last element (write zero value).
    				all_variables = all_variables[:len(all_variables)-1]   // Truncate slice.
    			}
    		}
    	}
    	//fmt.Println(all_variables)
    	return all_variables
    }

    func check_column(int_data [][]int, column int) []int {
    	all_variables := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    	//fmt.Println(all_variables)
    	//int_data
    	//i
    	for i := 0; i < 9; i++ {
    		data := int_data[i][column]
    		for j := 0; j < len(all_variables); j++ {
    			if data == all_variables[j] {
    				all_variables[j] = all_variables[len(all_variables)-1] // Copy last element to index i.
    				all_variables[len(all_variables)-1] = 0                // Erase last element (write zero value).
    				all_variables = all_variables[:len(all_variables)-1]   // Truncate slice.
    			}
    		}
    	}
    	//fmt.Println(all_variables)
    	return all_variables
    }

    func dots_to_zero(str string) string {
    	str1 := []rune(str)
    	for i, letter := range str1 {
    		if letter == '.' {
    			str1[i] = '0'
    		}
    	}
    	return string(str1)
    }

    func str_to_int(s [9]string) [][]int {
    	//s := []string{"096040001", "096040001"}
    	int_list_all := make([][]int, 9)

    	for number, str := range s {
    		int_list := make([]int, 9)
    		list := []rune(str)
    		for i := 0; i < 9; i++ {
    			int_list[i] = int(list[i]) - 48
    		}
    		int_list_all[number] = int_list
    	}
    	return int_list_all
    }
