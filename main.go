package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var Task_list []task 
var last_task_ID int = 0

func add_task(tid int,title, discription, add_datetime string){
	//adding task to the struct 
	new_task := task{
		task_ID: tid,
		task_title: title,
		task_add_datetime: add_datetime,
		task_description: discription,
		task_status: false,
	}

	Task_list = append(Task_list,new_task)

}


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

func update_sta(taskid int, sts bool){}

func del_task()  {}

func get_detail()  {}

func update_detail()  {}


func main() {
	//getting the user inputs to perform the operation  
	scanner := bufio.NewScanner(os.Stderr)
	fmt.Printf("Enter 1 to add new taks, 2 to update detail, 3 to update status and 4 to delete the task : ")
	scanner.Scan()
	operation := scanner.Text()
	//converting the user input in to int 
	opp_id, opp_err := strconv.Atoi(operation)
	if opp_err != nil {
		fmt.Printf("Error")
	}

	if (opp_id == 1){
		get_new_detail()
	}else{
		fmt.Printf("Enter the task ID : ")
		scanner.Scan()
		taskid := scanner.Text()

		//converting the user input in to int 
		tskid, err := strconv.Atoi(taskid)
		if err != nil {
			fmt.Printf("Error")
		}
		if(opp_id == 2){
			
			for tsk := 0 ; tsk < last_task_ID ; tsk++ {
				if(Task_list[tsk].task_ID == tskid){
					get_detail()
					break
				}else{
					fmt.Printf("Task ID not found")
				}
			}

		}else if(opp_id == 3){
			for tsk := 0 ; tsk < last_task_ID ; tsk++ {
				if(Task_list[tsk].task_ID == tskid){
					update_sta(tsk,true);
					break
				}else{
					fmt.Printf("Task ID not found")
				}
			}
		
		}else if(opp_id == 4){

			for tsk := 0 ; tsk < last_task_ID ; tsk++ {
				if(Task_list[tsk].task_ID == tsk){
					del_task()
					break
				}else{
					fmt.Printf("Task ID not found")
				}
			}
			
		}					
	}

}
