import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Task() {
    const [task, setTasks] = useState([]);
  
    const instance = axios.create({
      baseURL: 'http://localhost:8080',
      timeout: 1000,
    });
    delete instance.defaults.headers.common.Authorization

    // var taskID = this.props.match.params.id;
  
    useEffect(() => {
      instance.get('/tasks/' + taskID)
        .then(response => setTasks(response.data))
        .catch(error => console.error(error));
    }, []);
  
    return (
      <div class="flex-column flex-md-row p-4 gap-4 py-md-5 align-items-center justify-content-center">
            <div class="d-flex gap-2 w-75 justify-content-between">
                <h6 class="mb-0">Title: {task.title}</h6>
                <p class="mb-0 opacity-75">Description: {task.description}</p>
                <p class="mb-0 opacity-75">Status: {task.status}</p>
            </div>
      </div>
    );
  }

export default Task