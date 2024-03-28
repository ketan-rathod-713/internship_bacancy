import React, { useEffect, useRef, useState } from "react";
import { Manager } from "socket.io-client";

const Game = () => {
  let socketRef = useRef();
  const [userInfo, setUserInfo] = useState({
    Id: "user",
    Name: "Ketan Rathod",
  });
  const [playerType, setPlayerType] = useState(null); // 1 for room creater, 2 for room joiner
  const [gamePhase, setGamePhase] = useState({
    room: true,
    game: false,
    result: false,
  });
  const [roomError, setRoomError] = useState(null);
  const [currentGameStatus, setCurrentGameStatus] = useState({
    selection: true,
    waiting: false,
    result: false,
  });
  const [currentGameResult, setcurrentGameResult] = useState({
    userSelection: "",
    opponentSelection: "",
  });
  const [gameStartsIn, setgameStartsIn] = useState(5);

  const [roomId, setRoomId] = useState("");
  const [iterations, setIterations] = useState(10);
  const [gameResult, setGameResult] = useState([])

  const handleRoomIdChange = (newValue) => {
    setRoomId(newValue);
  };
  const handleIterationChange = (newValue) => {
    setIterations(newValue);
  };

  const createRoom = () => {
    console.log("create room");
    setPlayerType(1);
    socketRef.current.emit("createRoom", { iterations: Number(iterations) });
    console.log(iterations);
  };

  const joinRoom = () => {
    setPlayerType(2);
    socketRef.current.emit("joinRoom", { roomId: roomId });
    console.log("joining to room", roomId);
  };
  const handleSelection = (selection) => {
    // given integer
    setIterations(iterations - 1);
    socketRef.current.emit("select", { selected: selection });
    setcurrentGameResult({ userSelection: selection });

    // change status from selection to waiting
    setCurrentGameStatus({ result: true, selection: false, waiting: true });
    console.log("selected", selection);
  };

  const restartGame = () => {
    setGamePhase({ game: false, room: true, result: false });
    setCurrentGameStatus({ result: false, selection: true });
    setRoomError("");
    setRoomId("");
  };

  useEffect(() => {
    if (socketRef.current) {
      socketRef.current.on("currentGameResult", (data) => {
        console.log("game result", data);

        // for final game result
        if (data.data.isGameEnded) {
          // then show result
          setGameResult(data.data.gameResult)
          setGamePhase({ result: true, game: false, room: false });
        } else {
          // For current game result
          if (playerType == 1) {
            // if i am room creater then my selection would be player 1
            console.log(playerType);
            setcurrentGameResult({
              userSelection: data.data.player1Selection,
              opponentSelection: data.data.player2Selection,
            });
          } else if (playerType == 2) {
            console.log(playerType);
            setcurrentGameResult({
              userSelection: data.data.player2Selection,
              opponentSelection: data.data.player1Selection,
            });
          }

          setCurrentGameStatus({
            result: true,
            selection: false,
            waiting: false,
          });

          let interval = setInterval(() => {
            setgameStartsIn((prev) => prev - 1);
          }, 1000);

          setTimeout(() => {
            // erase previous result and all that state
            // show timer also
            clearInterval(interval);
            setgameStartsIn(5);
            setCurrentGameStatus({
              result: false,
              selection: true,
              waiting: false,
            });
          }, 5000);
        }
      });
    }
  }, [playerType]);

  useEffect(() => {
    //
    const manager = new Manager("ws://192.168.7.38:8080/", {
      reconnectionDelayMax: 10000,
      transports: ["websocket"],
      query: {
        data: "234",
      },
    });

    const socket = manager.socket("/", {
      auth: {
        token: "124",
      },
    });
    socketRef.current = socket;

    socket.on("connect", () => {
      console.log("Connection established");

      // if connection is closing then reset everything and mark connection closed variable
      setCurrentGameStatus({ result: false, selection: true, waiting: false });
      setGamePhase({ room: true, game: false, result: false });
      setRoomId("");
      setRoomError("");
      // socket.emit("game", {})
    });

    socket.on("connect_error", (err) => {
      console.log("Connection error", err);
    });

    socket.on("error", (err) => {
      console.error("Connection error", err);
    });

    socket.on("disconnect", () => {
      // if connection is closing then reset everything and mark connection closed variable
      setCurrentGameStatus({ result: false, selection: true, waiting: false });
      setGamePhase({ room: true, game: false, result: false });
      setRoomId("");
      setRoomError("");

      console.log("Connection Closed");
    });

    // for room creater
    socket.on("roomCreated", (data) => {
      console.log("room createds");
      console.log(data);
      setRoomId(data.data.roomId);
    });

    // for room joiner
    socket.on("joinedRoom", (data) => {
      console.log("room joined");
      console.log(data);

      if (data.status == "success") {
        setGamePhase({ game: true, room: false, result: false });
        setIterations(data.data.iterations);
      } else {
        alert("failed to join room");

        // todo WORK ON SHOWING ERROR
        setRoomError("failed to join room");
      }
    });

    socket.on("userDisconnected", (data) => {
      console.log("user disconnected");
      setGamePhase({ game: false, room: false, result: true });
    });
    return () => {};
  }, []);

  return (
    <div className="gamepage">
      <h1>Rock Paper Scissors</h1>
      <h1>{iterations}</h1>
      <div className="gamepage-main">
        {gamePhase.room && (
          <div className="card">
            <div className="create-join-room">
              <button onClick={createRoom}>Create Room</button>
              <button onClick={joinRoom}>Join Room</button>
            </div>
            <div className="roomIdInput">
              <div>
                <label htmlFor="input"> Number of iterations</label>
                <input
                  type="number"
                  onChange={(e) => handleIterationChange(e.target.value)}
                  value={iterations}
                  placeholder="eg. 10"
                />
              </div>
              <div>
                <label htmlFor="input"> Room Id</label>

                <input
                  type="text"
                  onChange={(e) => handleRoomIdChange(e.target.value)}
                  value={roomId}
                  placeholder="eg. 756849"
                />
              </div>
            </div>
          </div>
        )}

        {gamePhase.game && (
          <div className="gamephase">
            <div className="players-info">
              <div>Ketan</div>
              <div>VS</div>
              <div>Aman</div>
            </div>

            <div className="main">
              <div className="game-info">
                {currentGameStatus.selection && "Select Any One"}
                {currentGameStatus.result && "Game Result"}
              </div>
              {currentGameStatus.selection && (
                <div className="selection">
                  <div onClick={() => handleSelection("rock")}>Rock</div>
                  <div onClick={() => handleSelection("paper")}>Paper</div>
                  <div onClick={() => handleSelection("scissors")}>
                    Scissors
                  </div>
                </div>
              )}
              {currentGameStatus.waiting && <div>Waiting Bro Select Fast</div>}
              {currentGameStatus.result && !currentGameStatus.waiting && (
                <div className="gameStartsInTimer">
                  <h1>{gameStartsIn}</h1>
                </div>
              )}
              {currentGameStatus.result && (
                <div className="currentGameResult">
                  {/* if not waiting then show result */}
                  <div>
                    <div className="title">You Selected</div>
                    <div className="selection">
                      {currentGameResult.userSelection}
                    </div>
                  </div>

                  <div>
                    <div className="title">Aman Selected</div>
                    <div className="selection">
                      {currentGameStatus.waiting
                        ? "Waiting For Opponent To Choose..."
                        : currentGameResult.opponentSelection}
                    </div>
                  </div>
                </div>
              )}
            </div>
          </div>
        )}

        {gamePhase.result && (
          <div className="gamephaseresult">
            <div className="replaygame">
              <button>Request Replay Game</button>
              <button onClick={restartGame}>Restart Game</button>
            </div>
            <div className="result-table">
              <table>
                <thead>
                  <tr>
                    <th>Iteration</th>
                    <th>Player1Selection</th>
                    <th>Player2Selection</th>
                    <th>Winner</th>
                  </tr>
                </thead>
                <tbody>
                {gameResult.gamestate.map((result) => <tr>
                    <td>{ 1}</td>
                    <td>{result.player1Selected}</td>
                    <td>{result.player2Selected}</td>
                    <td>Winner</td>
                  </tr>)}
                </tbody>
              </table>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Game;
