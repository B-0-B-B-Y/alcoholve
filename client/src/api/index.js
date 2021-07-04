export const createNewGame = async () => {
  try {
    const response = await fetch('http://localhost:8080/new', {
      method: 'POST',
      body: JSON.stringify({
        playerNames: ["test1", "test2"],
        alcohol: "vodka",
        threshold: "1000",
        questionAmount: 12
      })
    })
    
    console.log(response)
  } catch (error) {
    console.error(error)
  }
}