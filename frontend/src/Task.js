import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router';
import { Link } from 'react-router-dom'
import axios from 'axios';

export default function Task() {
    const [task, setTask] = useState([]);
    const [error, setError] = useState([]);
  
    const instance = axios.create({
      baseURL: 'http://localhost:8080',
      timeout: 1000,
    });
    delete instance.defaults.headers.common.Authorization

    const params = useParams()

    var url = '/tasks/' + params.taskID
  
    useEffect(() => {
      instance.get(url)
        .then(response => setTask(response.data))
        .catch(error => setError(error));
    }, []);
  
    return (
      <div class="flex-column flex-md-row p-4 gap-4 py-md-5 align-items-center justify-content-center">
        <h3 class="mb-0">Title: {task.title}</h3>
        <h3 class="mb-0">Description: {task.description}</h3>
        <h3 class="mb-0">Status: {task.status}</h3>
        <p class="mb-0 opacity=50">Created: {task.created_at}</p>
        <p class="mb-0 opacity=50">Updated: {task.updated_at}</p>
        <Link class="btn btn-primary" to="/">Back</Link>
      </div>
    );
  }