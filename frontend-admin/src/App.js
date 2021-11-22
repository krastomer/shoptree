import './App.css';
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Home from './views/admin/Home';

function App() {
  return (
    <div>
      <BrowserRouter>
        <Switch>
          <Route path="/"><Home /></Route>
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
