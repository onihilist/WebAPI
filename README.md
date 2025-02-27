# Run the project
If it's the first time : 

    - $ make setup-env
    - $ make up

If you want to restart to apply Golang changes :

    - $ make restartapp

If you want to restart to apply MariaDB changes :

    - $ make restart

You must wait approx 50 seconds for getting the app operational on http://localhost:8080.
If you want to know when the app is ready, you can check docker logs : 

    - $ docker logs --tail 1000 -f <container_id>

If you want more information about make commands : 
    
    - $ make help

## Troubleshooting

If you start the app and go that error : 

    ```
        [2025-01-24|13:35:56.594089][FATAL] -> [MariaDB] - dial tcp 172.18.0.3:3306: connect: connection refused
            panic: runtime error: invalid memory address or nil pointer dereference
            panic: runtime error: invalid memory address or nil pointer dereference
            [signal SIGSEGV: segmentation violation code=0x1 addr=0x61 pc=0x52944e]
    ```

Just restart your app container, the app start a little too early.

## TODO

- [ ] Databases 
    - [ ] Add enable parameter in each database env file 
    - [ ] Make module for handle PostgreSQL
- [ ] Templates
    - [ ] Rework template User to send just the object
