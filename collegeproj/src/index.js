import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter , Routes, Route, } from "react-router-dom";
import './index.css';
import App from './App';
import Home from "./routes/home";
import Forum from "./routes/forum"
import TopicForm from './routes/topicForm';
import reportWebVitals from './reportWebVitals';
import Topic from './routes/topic';
import TopicList from './components/TopicList';


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
    <Routes>
      <Route path="/" element={<App />}>
        <Route path="home" element={<Home />} />
        <Route path="forum" element={<Forum />} >
          <Route path="all" element={<TopicList />} />
          <Route path="newtopic" element={<TopicForm />} />
          <Route path=":Id" element={<Topic />} />
        </Route>
      </Route>
    </Routes>
    </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
