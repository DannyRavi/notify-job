**README**

# **graph-co code exam**

This repository contains two Golang applications:

1. **File Writer:(Mario)** A simple application that writes a file using Goroutines. that name is **Mario**.
2. **File Monitor:(Luigi)** A more complex application that monitors a directory for file changes using Fanotify. that name is **Luigi**.

## **Software Requirements:**

  * **Operating System:**
      * Linux (e.g., Ubuntu, Debian, CentOS)
      * macOS
  * **Golang:**
      * Version 1.21 or later is recommended.
  * **makefile (for build, test, run):**
    ```sh
    sudo apt-get update
    sudo apt-get -y install make
    ```

## **File Writer (Mario)**

* **Functionality:**
    - Creates a new file or appends to an existing file.
    - Writes data to the file using multiple Goroutines for concurrent writes by Fan-in and Fan-out method.
    - Configurable number of Goroutines for parallel writes.

* **Usage: (makefile)**
    1. extract zip file and going to folder.
    2. clean file: `make clean`
    3. Build the application: `make build`
    4. Run the application: `make mario`

* **Usage: (manual)**
    1. extract zip file and going to folder.
    2. change path: `cd /cmd/mario/`
    3. Run the application: `go run . cli  -d /tmp/ox.csv -r 1000`

* **Flags: -d -r -v**
    1. `-d` for insert path of file -> /tmp/ox.csv
    2. `-r` number of goroutines -> between 1 to 10000
    3. `-v` verbose log mode 

* **sample of mario output**

```sh
# go run . cli -v  -d /tmp/ox5.csv -r 5000 

 ___      ___       __        _______    __      ______   
|"  \    /"  |     /""\      /"      \  |" \    /    " \  
 \   \  //   |    /    \    |:        | ||  |  // ____  \ 
 /\\  \/.    |   /' /\  \   |_____/   ) |:  | /  /    ) :)
|: \.        |  //  __'  \   //      /  |.  |(: (____/ // 
|.  \    /:  | /   /  \\  \ |:  __   \  /\  |\\        /  
|___|\__/|___|(___/    \___)|__|  \___)(__\_|_)\"_____/   
INFO[2024-11-20T02:21:19+03:30] number of goroutines: 5000                   
INFO[2024-11-20T02:21:19+03:30] output file path: /tmp/ox5.csv               
DEBU[2024-11-20T02:21:19+03:30] verbose mode: true                           
DEBU[2024-11-20T02:21:19+03:30] File /tmp/ox5.csv removed successfully.      
INFO[2024-11-20T02:21:19+03:30] number of Goroutines:5000                    
INFO[2024-11-20T02:21:47+03:30] Data written to successfully!                
INFO[2024-11-20T02:21:47+03:30] Execution time: 27.931528102s 
```

## **File Monitor (Luigi)**

* **Functionality:**
    - Monitors a specified directory for file creation events.
    - Leverages the `Fanotify` system call for efficient and low-overhead file system monitoring.
    - Logs file events to the console.
    - find `PID` of process that has been create a file
    - calculate `hash256` of file

* **Usage: (makefile)**
    1. extract zip file and going to folder.
    2. clean file: `make clean`
    3. Build the application: `make build`
    4. Run the application: `make luigi`

* **Usage: (manual)**
    0. noticed: you must run File Writer (Mario) app before to execute this app.
    1. extract zip file and going to folder.
    2. change path: `cd /cmd/luigi/`
    3. switch to root user: `sudo su`
    4. Run the application: `go run . cli -v -d /tmp/`


### **Running Both Applications Together**

* **Usage: (makefile)**
    1. extract zip file and going to folder.
    2. clean file: `make clean`
    3. Build the application: `make build`
    4. Run the application: `make run`

* **Usage: (manual)**
    1. Start the `mario` application.
    2. Start the `luigi` application, specifying the directory to monitor.

* **Running test Applications and Logs**
    1. enable Flag `-v` verbose log mode 
    2. for Running test for each App you should going to `cd /cmd/mario` or `cd /cmd/luigi` and run Test `go test`
    3. for run auto test you should use in mode user root `sudo su` and then run `make testapp`.

```sh
(sudo mode)
# go run . cli -v -d /tmp/


$$\                $$\           $$\ 
$$ |               \__|          \__|
$$ |     $$\   $$\ $$\  $$$$$$\  $$\ 
$$ |     $$ |  $$ |$$ |$$  __$$\ $$ |
$$ |     $$ |  $$ |$$ |$$ /  $$ |$$ |
$$ |     $$ |  $$ |$$ |$$ |  $$ |$$ |
$$$$$$$$\\$$$$$$  |$$ |\$$$$$$$ |$$ |
\________|\______/ \__| \____$$ |\__|
                       $$\   $$ |    
                       \$$$$$$  |    
                        \______/     
INFO[2024-11-20T02:01:32+03:30] Running as root...                           
DEBU[2024-11-20T02:01:32+03:30] verbose mode: true                           
DEBU[2024-11-20T02:01:32+03:30] luigi input dir: /tmp/                       
INFO[2024-11-20T02:01:56+03:30] File path: /tmp/aa.txt                       
INFO[2024-11-20T02:01:56+03:30] PID: 8871                                    
DEBU[2024-11-20T02:01:56+03:30] file created? true                           
INFO[2024-11-20T02:01:56+03:30] SHA256 hash: ead4c9dd2a3feb629119aa91972c2af0c68aeaa97e1606aaa619ccd57ce0491e 
DEBU[2024-11-20T02:01:56+03:30] File removed successfully.                   
DEBU[2024-11-20T02:01:56+03:30] /tmp/                                        
INFO[2024-11-20T02:01:56+03:30] Execution time:23.495221757s         

```

**Key Points**

* **Goroutines:** The `file_writer(mario)` utilizes Goroutines to achieve concurrent file writes, improving performance.
* **Fanotify:** The `file_monitor(luigi)` employs Fanotify to efficiently monitor file system events, providing real-time notifications.
* **Error Handling:** Both applications incorporate error handling mechanisms to gracefully handle potential issues.
* **Configuration:** The `file_writer(mario)` can be configured to adjust the number of Goroutines for parallel writes.

* **log and test:** Both applications incorporate log and test unit.

#### **Additional Considerations**

* **Fanotify Limitations:** Fanotify might have limitations or platform-specific behaviors, so it's essential to consider these factors when deploying in production environments.
* **Security:** Ensure proper file permissions and security measures to protect sensitive data.
* **Performance Optimization:** Experiment with different configurations and optimization techniques to fine-tune performance.

#### **Future Improvements**

* **docker or podman:** Docker and Podman are powerful tools for containerizing applications, providing numerous benefits when running Golang applicationsm likes: Consistency and Portability, Isolation and Security, Rapid Deployment and Scaling. but I have not enough time it use it!
* **Metric and monitoring tools:** Metric and monitoring tools operations for further accuracy gains.
* **Advanced Fanotify Features:** Explore advanced Fanotify features like mark events and user namespaces.

By combining Goroutines and Fanotify, these applications provide a robust and efficient solution for file operations and monitoring in Golang.


**Benchmarking**

compare number of goroutines vs execute time.

Here is a table that compares the number of goroutines vs total execution time:

| Number of Goroutines | Total Execution Time (seconds) |
|---|---|
| 100  | 152.97  |
| 500  | 31.09  |
| 1000  | 20.58  |
| 5000  | 27.93  |
| 10000  | 44.32  |

As the number of goroutines increases, the total execution time generally decreases. This is because Goroutines allow for concurrent execution of tasks, which can improve performance for CPU-bound tasks. However, there are also overhead costs associated with creating and managing Goroutines, so the optimal number of Goroutines will vary depending on the specific task and hardware.

