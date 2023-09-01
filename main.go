package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

)

type task struct {
	task_ID int ;
	task_title string ; 
	task_add_datetime string;
	task_description string ;
	task_status bool;
}

var Task_list []task
var last_task_ID int = 0

func get_new_detail(){
	last_task_ID ++
	//getting  the current data and time 
	task_add_time := time.Now() 
	// creating new scanner to get the user inputs
	scanner := bufio.NewScanner(os.Stdin)

	//getting the task title
	fmt.Printf("Enter the task title : ")
	scanner.Scan()
	task_title := scanner.Text()
	
	// gettinf the task discription 
	fmt.Printf("Enter the task Discription : ")
	scanner.Scan()
	taskdis := scanner.Text()

	add_task(last_task_ID,task_title,taskdis,task_add_time.String())
}

func main() {
	// Getting the user inputs to perform the operation
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter 1 to add new task, 2 to update detail, 3 to update status and 4 to delete the task: ")
	scanner.Scan()
	operation := scanner.Text()
	// Converting the user input into an int
	opp_id, opp_err := strconv.Atoi(operation)
	if opp_err != nil {
		fmt.Printf("Error: %v", opp_err)
	}

	if opp_id == 1 {
		get_new_detail()
	} else {
		fmt.Printf("Enter the task ID: ")
		scanner.Scan()
		taskid := scanner.Text()

		// Converting the user input into an int
		tskid, err := strconv.Atoi(taskid)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		if opp_id == 2 {

			for tsk := 0; tsk < last_task_ID; tsk++ {
				if Task_list[tsk].task_ID == tskid {
					get_detail(tskid)
					break
				} else {
					fmt.Println("Task ID not found")
				}
			}

		} else if opp_id == 3 {
			for tsk := 0; tsk < last_task_ID; tsk++ {
				if Task_list[tsk].task_ID == tskid {
					fmt.Printf("Enter 1 to mark as complete: ")
					scanner.Scan()
					stat := scanner.Text()
					ss_id, ss_err := strconv.Atoi(stat)
					if ss_err != nil {
						fmt.Printf("Error: %v", ss_err)
					} else if ss_id == 1 {
						update_sta(tskid, true)
					} else {
						update_sta(tskid, false)
					}

					break
				} else {
					fmt.Println("Task ID not found")
				}
			}

		} else if opp_id == 4 {

			for tsk := 0; tsk < last_task_ID; tsk++ {
				if Task_list[tsk].task_ID == tskid {
					del_task(tskid)
					break
				} else {
					fmt.Println("Task ID not found")
				}
			}

		}
	}

}

func add_task(tid int, title, discription, add_datetime string) {
	// Adding task to the struct
	new_task := task{
		task_ID:           tid,
		task_title:        title,
		task_add_datetime: add_datetime,
		task_description:  discription,
		task_status:       false,
	}

	Task_list = append(Task_list, new_task)
	fmt.Println("Task added successfully !!")
}

func update_sta(taskid int, sts bool) {
	for index := 0; index < last_task_ID; index++ {
		if Task_list[index].task_ID == taskid {

			Task_list[index].task_status = sts
			fmt.Println("Task status updated !!")
			break
		}
	}
}

func del_task(id int) {
	for index := 0; index < last_task_ID; index++ {
		if Task_list[index].task_ID == id {
			Task_list = append(Task_list[:index], Task_list[index+1:]...)
			fmt.Println("Task deleted successfully !!")
			break
		}
	}
}

func get_detail(tid int) {
	scanner := bufio.NewScanner(os.Stdin)
	for tsk := 0; tsk < last_task_ID; tsk++ {
		if Task_list[tsk].task_ID == tid {
			fmt.Print("Enter the task Title: ")
			scanner.Scan()
			ttital := scanner.Text()
			fmt.Print("Enter the task Description: ")
			scanner.Scan()
			tdis := scanner.Text()
			tas_di := Task_list[tsk].task_description
			tast := Task_list[tsk].task_title
			if ttital != "" {
				update_detail(ttital, tas_di, tid)
			} else if tdis != "" {
				update_detail(tast, tdis, tid)
			} else if ttital != "" && tdis != "" {
				update_detail(ttital, tdis, tid)
			} else {
				fmt.Println("Error")
			}
			break
		}
	}
}

func update_detail(tittle string, discr string, id int) {
	for index := 0; index < last_task_ID; index++ {

		if Task_list[index].task_ID == id {
			Task_list[index].task_title = tittle
			Task_list[index].task_description = discr
			fmt.Println("Task updated successfully !!")
			break
		}
	}
}
