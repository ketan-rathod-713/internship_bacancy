import './App.css';
import { BrowserRouter as Router, Route, BrowserRouter, RouterProvider, createBrowserRouter, Link } from 'react-router-dom';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Profile from './pages/Profile';
import Game from './pages/Game';
import LoginForm from './pages/Login';

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <div>
        <h1>Hello World</h1>
        <Link to="game">Play Game</Link>
        <Link to="profile">Profile</Link>
        <Link to="login">Login</Link>
        <Link to="signup">Signup</Link>
      </div>
    ),
  },
  {
    path: "game",
    element: <Game/>,
  },
  {
    path: "profile",
    element: <Profile/>,
  },
  {
    path: "login",
    element: <LoginForm/>,
  },
  {
    path: "/signup",
    element: <Signup/>,
  },
]);

function App() {
  return (
    <RouterProvider router={router}/>
  );
}

export default App;
