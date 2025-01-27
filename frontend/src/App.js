import React from 'react';
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Home from './Home.js'
import Task from './Task.js'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Home/>
  },
  {
    path: '/task/:taskID',
    element: <Task/>
  }
]);
function App() {
  return (
    <RouterProvider router={router} />
  );
}

export default App;
