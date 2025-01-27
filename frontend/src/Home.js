import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom'
import axios from 'axios';

export default function Home() {
  const [tasks, setTasks] = useState([]);

  const instance = axios.create({
    baseURL: 'http://localhost:8080',
    timeout: 1000,
  });
  delete instance.defaults.headers.common.Authorization

  useEffect(() => {
    instance.get('/tasks')
      .then(response => setTasks(response.data))
      .catch(error => console.error(error));
  }, []);

  return (
    <div class="flex-column flex-md-row p-4 gap-4 py-md-5 align-items-center justify-content-center">
      <h1 class="mb-0">Task List</h1><br/>
      <div class="list-group">
        {tasks.map(task => (
          <Link to={`/task/${task.id}`} class="list-group-item list-group-item-action w-75 gap-3 py-3" aria-current="true">
            <div class="d-flex gap-2 w-75 justify-content-between">
              <div>
                <h6 class="mb-0">{task.title}</h6>
                <p class="mb-0 opacity-75">{task.description}</p>
                <p class="mb-0 opacity-75">{task.status}</p>
              </div>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
}