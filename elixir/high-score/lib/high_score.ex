defmodule HighScore do
  def new() do
    %{}
  end

  @score 0
  def add_player(scores, name, score \\ @score) do
    Map.put(scores, name, score)
  end

  def remove_player(scores, name) do
    Map.delete(scores, name)
  end

  def reset_score(scores, name) do
    Map.put(scores, name, @score)
  end

  def update_score(scores, name, score \\ @score) do
    Map.update(scores, name, score,
      fn existing_score -> existing_score + score end
    )
  end

  def get_players(scores) do
    Map.keys(scores)
  end
end
