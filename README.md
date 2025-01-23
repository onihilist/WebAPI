# Run the project
If it's the first time : 

    - $ make up

If you want to restart to apply golang changes :

    - $ make down
    - $ make clean
    - $ make up

You must wait approx 50 seconds for getting the app operational on http://localhost:8080.
If you want to know when the app is ready, you can check docker logs : 

    - docker logs --tail 1000 -f <container_id>
