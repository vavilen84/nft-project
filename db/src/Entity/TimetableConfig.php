<?php

namespace App\Entity;

use App\Repository\TimetableConfigRepository;
use Doctrine\ORM\Mapping as ORM;

/**
 * @ORM\Entity(repositoryClass=TimetableConfigRepository::class)
 */
class TimetableConfig
{
    /**
     * @ORM\Id
     * @ORM\GeneratedValue
     * @ORM\Column(type="integer")
     */
    private $id;

    /**
     * @ORM\OneToOne(targetEntity=User::class, inversedBy="timetableConfig", cascade={"persist", "remove"})
     * @ORM\JoinColumn(nullable=false)
     */
    private $tutor;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $monday;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $tuesday;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $wednesday;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $thursday;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $friday;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $saturday;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $sunday;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $mondayStartHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $mondayStopHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $tuesdayStartHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $tuesdayStopHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $wednesdayStartHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $wednesdayStopHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $thursdayStartHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $thursdayStopHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $fridayStartHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $fridayStopHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $saturdayStartHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $saturdayStopHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $sundayStartHour;

    /**
     * @ORM\Column(type="integer", nullable=true)
     */
    private $sundayStopHour;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getTutor(): ?User
    {
        return $this->tutor;
    }

    public function setTutor(User $tutor): self
    {
        $this->tutor = $tutor;

        return $this;
    }

    public function isMonday(): ?bool
    {
        return $this->monday;
    }

    public function setMonday(?bool $monday): self
    {
        $this->monday = $monday;

        return $this;
    }

    public function isTuesday(): ?bool
    {
        return $this->tuesday;
    }

    public function setTuesday(?bool $tuesday): self
    {
        $this->tuesday = $tuesday;

        return $this;
    }

    public function isWednesday(): ?bool
    {
        return $this->wednesday;
    }

    public function setWednesday(?bool $wednesday): self
    {
        $this->wednesday = $wednesday;

        return $this;
    }

    public function isThursday(): ?bool
    {
        return $this->thursday;
    }

    public function setThursday(?bool $thursday): self
    {
        $this->thursday = $thursday;

        return $this;
    }

    public function isFriday(): ?bool
    {
        return $this->friday;
    }

    public function setFriday(?bool $friday): self
    {
        $this->friday = $friday;

        return $this;
    }

    public function isSaturday(): ?bool
    {
        return $this->saturday;
    }

    public function setSaturday(?bool $saturday): self
    {
        $this->saturday = $saturday;

        return $this;
    }

    public function isSunday(): ?bool
    {
        return $this->sunday;
    }

    public function setSunday(?bool $sunday): self
    {
        $this->sunday = $sunday;

        return $this;
    }

    public function getMondayStartHour(): ?int
    {
        return $this->mondayStartHour;
    }

    public function setMondayStartHour(?int $mondayStartHour): self
    {
        $this->mondayStartHour = $mondayStartHour;

        return $this;
    }

    public function getMondayStopHour(): ?int
    {
        return $this->mondayStopHour;
    }

    public function setMondayStopHour(?int $mondayStopHour): self
    {
        $this->mondayStopHour = $mondayStopHour;

        return $this;
    }

    public function getTuesdayStartHour(): ?int
    {
        return $this->tuesdayStartHour;
    }

    public function setTuesdayStartHour(?int $tuesdayStartHour): self
    {
        $this->tuesdayStartHour = $tuesdayStartHour;

        return $this;
    }

    public function getTuesdayStopHour(): ?int
    {
        return $this->tuesdayStopHour;
    }

    public function setTuesdayStopHour(?int $tuesdayStopHour): self
    {
        $this->tuesdayStopHour = $tuesdayStopHour;

        return $this;
    }

    public function getWednesdayStartHour(): ?int
    {
        return $this->wednesdayStartHour;
    }

    public function setWednesdayStartHour(?int $wednesdayStartHour): self
    {
        $this->wednesdayStartHour = $wednesdayStartHour;

        return $this;
    }

    public function getWednesdayStopHour(): ?int
    {
        return $this->wednesdayStopHour;
    }

    public function setWednesdayStopHour(?int $wednesdayStopHour): self
    {
        $this->wednesdayStopHour = $wednesdayStopHour;

        return $this;
    }

    public function getThursdayStartHour(): ?int
    {
        return $this->thursdayStartHour;
    }

    public function setThursdayStartHour(?int $thursdayStartHour): self
    {
        $this->thursdayStartHour = $thursdayStartHour;

        return $this;
    }

    public function getThursdayStopHour(): ?int
    {
        return $this->thursdayStopHour;
    }

    public function setThursdayStopHour(?int $thursdayStopHour): self
    {
        $this->thursdayStopHour = $thursdayStopHour;

        return $this;
    }

    public function getFridayStartHour(): ?int
    {
        return $this->fridayStartHour;
    }

    public function setFridayStartHour(?int $fridayStartHour): self
    {
        $this->fridayStartHour = $fridayStartHour;

        return $this;
    }

    public function getFridayStopHour(): ?int
    {
        return $this->fridayStopHour;
    }

    public function setFridayStopHour(?int $fridayStopHour): self
    {
        $this->fridayStopHour = $fridayStopHour;

        return $this;
    }

    public function getSaturdayStartHour(): ?int
    {
        return $this->saturdayStartHour;
    }

    public function setSaturdayStartHour(?int $saturdayStartHour): self
    {
        $this->saturdayStartHour = $saturdayStartHour;

        return $this;
    }

    public function getSaturdayStopHour(): ?int
    {
        return $this->saturdayStopHour;
    }

    public function setSaturdayStopHour(?int $saturdayStopHour): self
    {
        $this->saturdayStopHour = $saturdayStopHour;

        return $this;
    }

    public function getSundayStartHour(): ?int
    {
        return $this->sundayStartHour;
    }

    public function setSundayStartHour(?int $sundayStartHour): self
    {
        $this->sundayStartHour = $sundayStartHour;

        return $this;
    }

    public function getSundayStopHour(): ?int
    {
        return $this->sundayStopHour;
    }

    public function setSundayStopHour(?int $sundayStopHour): self
    {
        $this->sundayStopHour = $sundayStopHour;

        return $this;
    }
}
