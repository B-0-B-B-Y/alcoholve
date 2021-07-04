import '../styles/App.css';

export const App = () => {
  return (
    <div className='app'>
      <h1 className='title'>Welcome to Alcoholve</h1>

      <div className='menu'>
        <button>Create a new game</button>
        <button>Join a game</button>
        <button>Help</button>
      </div>
    </div>
  );
}
