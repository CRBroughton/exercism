defmodule Lasagna do
  # Please define the 'expected_minutes_in_oven/0' function
  @spec expected_minutes_in_oven() :: number()
  def expected_minutes_in_oven, do: 40

  # Please define the 'remaining_minutes_in_oven/1' function
  @spec remaining_minutes_in_oven(number()) :: number()
  def remaining_minutes_in_oven(minutes), do: expected_minutes_in_oven() - minutes

  # Please define the 'preparation_time_in_minutes/1' function
  @spec preparation_time_in_minutes(number()) :: number()
  def preparation_time_in_minutes(layers), do: layers * 2

  @spec total_time_in_minutes(number(), number()) :: number()
  # Please define the 'total_time_in_minutes/2' function
  def total_time_in_minutes(layers, minutes), do: preparation_time_in_minutes(layers) + minutes

  @spec alarm() :: String.t()
  # Please define the 'alarm/0' function
  def alarm, do: "Ding!"
end
