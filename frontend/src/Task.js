import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router';
import axios from 'axios';

export default function Task() {
    const [task, setTask] = useState([]);
  
    const instance = axios.create({
      baseURL: 'http://localhost:8080',
      timeout: 1000,
    });
    delete instance.defaults.headers.common.Authorization

    const params = useParams()
    console.log(params)

    var url = '/tasks/' + params.taskID
    console.log(url)
  
    useEffect(() => {
      instance.get(url)
        .then(response => setTask(response.data))
        .catch(error => console.error(error));
    }, []);

    console.log(task)
  
    return (
      <div class="flex-column flex-md-row p-4 gap-4 py-md-5 align-items-center justify-content-center">
                <h6 class="mb-0">Title: {task.title}</h6>
                <p class="mb-0 opacity-75">Description: {task.description}</p>
                <p class="mb-0 opacity-75">Status: {task.status}</p>
      </div>
    );
  }