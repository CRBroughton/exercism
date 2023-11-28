defmodule Darts do
  import :math, only: [sqrt: 1]
  @type position :: {number, number}

  @doc """
  Calculate the score of a single dart hitting a target
  """
  @spec score(position) :: integer
  def score({x, y}) do
    distance = sqrt(x * x + y * y)
    cond do
      distance <= 1 -> 10
      distance <= 5 -> 5
      distance <= 10 -> 10
      true -> 0
    end
  end
end
