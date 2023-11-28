<?php

class AnnalynsInfiltration
{
    public function canFastAttack($is_knight_awake)
    {
        return !$is_knight_awake;
    }

    public function canSpy(
        $is_knight_awake,
        $is_archer_awake,
        $is_prisoner_awake
    ) {
        $someoneIsAwake = $is_knight_awake || $is_archer_awake || $is_prisoner_awake;
        return $someoneIsAwake;
    }

    public function canSignal(
        $is_archer_awake,
        $is_prisoner_awake
    ) {
        $canRescuePrisoner = $is_archer_awake === false && $is_prisoner_awake;
        return $canRescuePrisoner;
    }

    public function canLiberate(
        $is_knight_awake,
        $is_archer_awake,
        $is_prisoner_awake,
        $is_dog_present
    ) {
        $guardsAreAsleep = !$is_archer_awake && !$is_knight_awake;
        return $is_dog_present && !$is_archer_awake || !$is_dog_present && $is_prisoner_awake && $guardsAreAsleep;
    }
}
