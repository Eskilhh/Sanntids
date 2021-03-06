package FSM

import (
	"../Driver"
	"../Network_main"
	"../Timer"
	"fmt"
	"time"
)

//scp -r Heis student@129.241.187.146:~/grr/
//ssh student@129.241.187.146

func Function_state_machine() {
	Arrived_chan := make(chan bool, 100000000)
	Order_chan := make(chan bool, 100000000)
	Set_timeout_chan := make(chan bool, 100000000)
	Set_timer_chan := make(chan bool, 100000000)

	go Network_main.Network_main(Order_chan)
	go Network_main.Order_compare_outer_list()
	go Network_main.Cost_function()
	go Driver.Lights_tracking()
	go Driver.Is_arrived(Arrived_chan, Set_timeout_chan)
	go Driver.Order_set_inner_order()
	go Driver.Order_set_outer_order()
	go Driver.Set_current_floor()
	go Driver.Register_button(Order_chan)
	go Driver.Elev_is_idle(Order_chan)
	go Timer.Timer(Set_timeout_chan, Set_timer_chan, Order_chan)

	go Driver.Print_queue()
	for {
		select {

		case <-Arrived_chan:
			Driver.Elev_set_motor_dir(Driver.DIRN_STOP)
			dir := Driver.Next_order()
			Driver.Elev_set_motor_dir(dir)
			Set_timer_chan <- true
			Driver.Elev_set_door_open_lamp(true)
		case <-Order_chan:
			dir := Driver.Next_order()
			Driver.Elev_set_motor_dir(dir)
		case <-Set_timeout_chan:
			Driver.Elev_set_motor_dir(Driver.DIRN_STOP)
			Driver.Elev_set_door_open_lamp(false)
		}
	}
}

func SLETT_DENNE() {
	fmt.Println("SLETT DEN DAA")
	time.Sleep(1 * time.Second)
}
