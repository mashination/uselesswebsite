import logo from './logo.svg';
import forumReducer from './store/reducers/forum'
import ReduxThunk from 'redux-thunk';
import { Provider } from 'react-redux';
import { createStore, combineReducers, applyMiddleware } from 'redux';
import './App.css';
import { Outlet, Link } from "react-router-dom";

const rootReducer = combineReducers({
  forum: forumReducer,
});

const store = createStore(rootReducer, applyMiddleware(ReduxThunk));

function App() {
  return (
    <Provider store = {store}>
    <div className="App">
      <div className='App-header'>
        <h2> The Useless Website</h2>
      </div>
      <div className='Menu'>
        <nav
          style={{
            borderBottom: "solid 1px",
            paddingBottom: "1rem",
          }}
        >
          <Link className='Home' style = {{color : '#FFF'}} to="/home">home</Link> |{" "}
          <Link className='Forum' style = {{color : '#FFF'}} to="/forum/all">forum</Link>
        </nav>
      </div>
      <Outlet/>
      
    </div>
    </Provider>
  );
}

export default App;
