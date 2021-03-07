import {BrowserRouter, Switch, Route} from 'react-router-dom'
import VideoPage from './components/videoPage'

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/login">
          <Login />
        </Route>
        <Route path="/videos">
          <VideoPage />
        </Route>
        <Route path="/">
          <VideoPage />
        </Route>
      </Switch>
    </BrowserRouter>
  );
}

function Login() {
  return 'Login'
}

export default App;
