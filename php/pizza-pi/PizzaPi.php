<?php

class PizzaPi
{
    public function calculateDoughRequirement(int $pizzas, int $persons)
    {
        return $pizzas * (($persons * 20) + 200);
    }

    public function calculateSauceRequirement(int $pizzas, int $can_volume)
    {
        return $pizzas * 125 / $can_volume;
    }

    public function calculateCheeseCubeCoverage(int $dimension, int|float $thickness, int $diameter)
    {
        return floor(($dimension ** 3) / ($thickness * M_PI * $diameter));
    }

    public function calculateLeftOverSlices(int $pizzas, int $persons)
    {
        return $pizzas * 8 % $persons;
    }
}
