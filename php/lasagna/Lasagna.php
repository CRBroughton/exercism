<?php

declare(strict_types=1);

class Lasagna
{
    // Please define the 'expectedCookTime()' function
    function expectedCookTime() 
    {
        return 40;
    }

    // Please define the 'remainingCookTime($elapsed_minutes)' function
    function remainingCookTime(int $minutes_in_oven)
    {
        return Lasagna::expectedCookTime() - $minutes_in_oven;
    }

    // Please define the 'totalPreparationTime($layers_to_prep)' function
    function totalPreparationTime(int $layers)
    {
        return $layers * 2;
    }

    // Please define the 'totalElapsedTime($layers_to_prep, $elapsed_minutes)' function
    function totalElapsedTime(int $layers, int $minutes_in_oven)
    {
        return Lasagna::totalPreparationTime($layers) + $minutes_in_oven;
    }

    // Please define the 'alarm()' function
    function alarm() {
        return "Ding!";
    }
}
