import './App.css';
import {Controls} from "./components/controls/Controls";
import {Map} from "./components/Map/Map";

function App() {
    return (
        <div className="App">
            <Map/>
            <Controls/>
        </div>
    );
}

export default App;
