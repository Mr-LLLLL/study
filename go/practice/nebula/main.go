package main

import (
	"fmt"
	"strings"
	"sync"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

const (
	address  = "127.0.0.1"
	port     = 9669
	username = "root"
	password = "nebula"
)

// Initialize logger
var log = nebula.DefaultLogger{}

func main() {
	hostAddress := nebula.HostAddress{Host: address, Port: port}
	hostList := []nebula.HostAddress{hostAddress}
	// Create configs for connection pool using default values
	testPoolConfig := nebula.GetDefaultConf()

	// Initialize connection pool
	pool, err := nebula.NewConnectionPool(hostList, testPoolConfig, log)

	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to initialize the connection pool, host: %s, port: %d, %s", address, port, err.Error()))
	}
	// Close all connections in the pool
	defer pool.Close()

	// Create session
	session, err := pool.GetSession(username, password)
	if err != nil {
		log.Fatal(fmt.Sprintf("Fail to create a new session from connection pool, username: %s, password: %s, %s",
			username, password, err.Error()))
	}
	// Release session and return connection back to connection pool
	defer session.Release()

	checkResultSet := func(prefix string, res *nebula.ResultSet) {
		if !res.IsSucceed() {
			log.Fatal(fmt.Sprintf("%s, ErrorCode: %v, ErrorMsg: %s", prefix, res.GetErrorCode(), res.GetErrorMsg()))
		}
	}

	{
		// Prepare the query
		createSchema := "CREATE SPACE IF NOT EXISTS test1(vid_type=int64); " +
			"USE test1;" +
			"CREATE TAG IF NOT EXISTS people1(name string not null default '', age int not null default 0, alias string, sex bool, birthday datetime);"

		// Execute a query
		resultSet, err := session.Execute(createSchema)

		if err != nil {
			fmt.Print(err.Error())
			return
		}
		checkResultSet(createSchema, resultSet)

		wg := sync.WaitGroup{}
		ch := make(chan struct{}, 100)
		for i := 0; i < 1000; i++ {
			i := i
			wg.Add(1)
			go func() {
				ch <- struct{}{}
				defer func() {
					wg.Done()
					<-ch
				}()
				start := 10000 * i
				end := 10000 * (i + 1)

				insert := strings.Builder{}
				insert.WriteString("insert vertex people1(name, age, alias, sex, birthday) VALUES ")
				for i := start; i < end; i++ {
					insert.WriteString(fmt.Sprintf(`%d:("%d", %d, "%d", %t, datetime()),`, i, i, i, i, i%2 == 1))
				}
				insert.WriteString(fmt.Sprintf(`%d:("%d", %d, "%d", %t, datetime());`, end, end, end, end, end%2 == 1))
				fmt.Println(insert.String())
				resultSet, err = session.Execute(insert.String())
				if err != nil {
					fmt.Print(err.Error())
					return
				}
				if !resultSet.IsSucceed() {
					log.Fatal(resultSet.GetErrorMsg())
				}
			}()
		}
		wg.Wait()
	}
	// {
	// 	// Prepare the query
	// 	createSchema := "USE test;" + "CREATE TAG IF NOT EXISTS group(name string, degree int, position string);"
	//
	// 	// Execute a query
	// 	resultSet, err := session.Execute(createSchema)
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 		return
	// 	}
	// 	checkResultSet(createSchema, resultSet)
	//
	// 	wg := sync.WaitGroup{}
	// 	ch := make(chan struct{}, 100)
	// 	for i := 1000; i < 2000; i++ {
	// 		i := i
	// 		wg.Add(1)
	// 		go func() {
	// 			ch <- struct{}{}
	// 			defer func() {
	// 				wg.Done()
	// 				<-ch
	// 			}()
	// 			start := 10000 * i
	// 			end := 10000 * (i + 1)
	//
	// 			insert := strings.Builder{}
	// 			insert.WriteString("insert vertex group(name, degree, position) VALUES ")
	// 			for i := start; i < end; i++ {
	// 				insert.WriteString(fmt.Sprintf(`%d:("%d", %d, "%d"),`, i, i, i, i))
	// 			}
	// 			insert.WriteString(fmt.Sprintf(`%d:("%d", %d, "%d");`, end, end, end, end))
	// 			fmt.Println(insert.String())
	// 			resultSet, err = session.Execute(insert.String())
	// 			if err != nil {
	// 				fmt.Print(err.Error())
	// 				return
	// 			}
	// 			if !resultSet.IsSucceed() {
	// 				log.Fatal(resultSet.GetErrorMsg())
	// 			}
	// 		}()
	// 	}
	// 	wg.Wait()
	// }
	{
		// Prepare the query
		createSchema := "CREATE SPACE IF NOT EXISTS test(vid_type=int64); " +
			"USE test;" +
			"CREATE EDGE IF NOT EXISTS like(likeness double);" +
			"CREATE EDGE IF NOT EXISTS join(join_datetime datetime);"

		// Execute a query
		resultSet, err := session.Execute(createSchema)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		checkResultSet(createSchema, resultSet)

		wg := sync.WaitGroup{}
		// ch := make(chan struct{}, 100)
		// for i := 0; i < 1000; i++ {
		// 	i := i
		// 	wg.Add(1)
		// 	go func() {
		// 		ch <- struct{}{}
		// 		defer func() {
		// 			wg.Done()
		// 			<-ch
		// 		}()
		// 		start := 10000 * i
		// 		end := 10000 * (i + 1)
		//
		// 		insert := strings.Builder{}
		// 		insert.WriteString("insert vertex people(name, age) VALUES ")
		// 		for i := start; i < end; i++ {
		// 			insert.WriteString(fmt.Sprintf(`%d:("%d", %d),`, i, i, i))
		// 		}
		// 		insert.WriteString(fmt.Sprintf(`%d:("%d", %d);`, end, end, end))
		// 		fmt.Println(insert.String())
		// 		resultSet, err = session.Execute(insert.String())
		// 		if err != nil {
		// 			fmt.Print(err.Error())
		// 			return
		// 		}
		// 		if !resultSet.IsSucceed() {
		// 			log.Fatal(resultSet.GetErrorMsg())
		// 		}
		// 	}()
		// }
		wg.Wait()
	}

	// {
	// 	// Prepare the query
	// 	createSchema := "CREATE SPACE IF NOT EXISTS basic_example_space(vid_type=FIXED_STRING(20)); " +
	// 		"USE basic_example_space;" +
	// 		"CREATE TAG IF NOT EXISTS person(name string, age int);" +
	// 		"CREATE EDGE IF NOT EXISTS like(likeness double)"
	//
	// 	// Execute a query
	// 	resultSet, err := session.Execute(createSchema)
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 		return
	// 	}
	// 	checkResultSet(createSchema, resultSet)
	// }
	// Drop space
	// {
	// 	query := "DROP SPACE IF EXISTS basic_example_space"
	// 	// Send query
	// 	resultSet, err := session.Execute(query)
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 		return
	// 	}
	// 	checkResultSet(query, resultSet)
	// }

	fmt.Print("\n")
	log.Info("Nebula Go Client Basic Example Finished")
}
