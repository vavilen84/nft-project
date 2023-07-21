<?php

namespace App\Entity;

use App\Repository\UserRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;
use Symfony\Component\Security\Core\User\PasswordAuthenticatedUserInterface;
use Symfony\Component\Security\Core\User\UserInterface;
use Symfony\Bridge\Doctrine\Validator\Constraints\UniqueEntity;

/**
 * @ORM\Entity(repositoryClass=UserRepository::class)
 * @ORM\Table(name="`user`")
 * @UniqueEntity("email")
 */
class User implements UserInterface, PasswordAuthenticatedUserInterface
{
    /**
     * @ORM\Id
     * @ORM\GeneratedValue
     * @ORM\Column(type="integer")
     */
    private $id;

    /**
     * @ORM\Column(type="string", length=180, unique=true)
     */
    private $email;

    /**
     * @ORM\Column(type="json")
     */
    private $roles = [];

    /**
     * @var string The hashed password
     * @ORM\Column(type="string")
     */
    private $password;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $firstName;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $lastName;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $telegram;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $skype;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $zoomLink;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $resetPasswordToken;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $emailVerificationHash;

    /**
     * @ORM\Column(type="datetime_immutable", nullable=true)
     */
    private $passwordResetHashExpiredAt;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $isEmailVerified;

    /**
     * @ORM\Column(type="boolean", nullable=true)
     */
    private $isTutorProfilePublished;

    /**
     * @ORM\Column(type="text", nullable=true)
     */
    private $textPresentation;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $videoPresentation;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $photo;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $countryOfBirth;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $currentCountry;

    /**
     * @ORM\Column(type="json", nullable=true)
     */
    private $subjects = [];

    /**
     * @ORM\Column(type="date")
     */
    private $birthday;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $timezone;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $gender;

    /**
     * @ORM\Column(type="datetime_immutable")
     */
    private $createdAt;

    /**
     * @ORM\Column(type="datetime_immutable")
     */
    private $updatedAt;

    /**
     * @ORM\Column(type="datetime_immutable", nullable=true)
     */
    private $deletedAt;

    /**
     * @ORM\OneToMany(targetEntity=Lesson::class, mappedBy="student")
     */
    private $lessons;

    /**
     * @ORM\Column(type="float", nullable=true)
     */
    private $tonRate;

    /**
     * @ORM\Column(type="json", nullable=true)
     */
    private $tutorLanguages = [];

    /**
     * @ORM\OneToMany(targetEntity=UserBalance::class, mappedBy="user", orphanRemoval=true)
     */
    private $userBalances;

    /**
     * @ORM\Column(type="string", length=255, nullable=true)
     */
    private $tonWalletAddress;

    /**
     * @ORM\OneToMany(targetEntity=Review::class, mappedBy="user")
     */
    private $reviews;

    /**
     * @ORM\OneToOne(targetEntity=TimetableConfig::class, mappedBy="tutor", cascade={"persist", "remove"})
     */
    private $timetableConfig;

    /**
     * @ORM\OneToMany(targetEntity=TimeslotReservation::class, mappedBy="tutor")
     */
    private $timeslotReservations;

    public function __construct()
    {
        $this->lessons = new ArrayCollection();
        $this->userBalances = new ArrayCollection();
        $this->reviews = new ArrayCollection();
        $this->timeslotReservations = new ArrayCollection();
    }

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getEmail(): ?string
    {
        return $this->email;
    }

    public function setEmail(string $email): self
    {
        $this->email = $email;

        return $this;
    }

    /**
     * A visual identifier that represents this user.
     *
     * @see UserInterface
     */
    public function getUserIdentifier(): string
    {
        return (string) $this->email;
    }

    /**
     * @deprecated since Symfony 5.3, use getUserIdentifier instead
     */
    public function getUsername(): string
    {
        return (string) $this->email;
    }

    /**
     * @see UserInterface
     */
    public function getRoles(): array
    {
        $roles = $this->roles;

        return array_unique($roles);
    }

    public function setRoles(array $roles): self
    {
        $this->roles = $roles;

        return $this;
    }

    /**
     * @see PasswordAuthenticatedUserInterface
     */
    public function getPassword(): string
    {
        return $this->password;
    }

    public function setPassword(string $password): self
    {
        $this->password = $password;

        return $this;
    }

    /**
     * Returning a salt is only needed, if you are not using a modern
     * hashing algorithm (e.g. bcrypt or sodium) in your security.yaml.
     *
     * @see UserInterface
     */
    public function getSalt(): ?string
    {
        return null;
    }

    /**
     * @see UserInterface
     */
    public function eraseCredentials()
    {
        // If you store any temporary, sensitive data on the user, clear it here
        // $this->plainPassword = null;
    }

    public function getFirstName(): ?string
    {
        return $this->firstName;
    }

    public function setFirstName(string $firstName): self
    {
        $this->firstName = $firstName;

        return $this;
    }

    public function getLastName(): ?string
    {
        return $this->lastName;
    }

    public function setLastName(string $lastName): self
    {
        $this->lastName = $lastName;

        return $this;
    }

    public function getTelegram(): ?string
    {
        return $this->telegram;
    }

    public function setTelegram(?string $telegram): self
    {
        $this->telegram = $telegram;

        return $this;
    }

    public function getSkype(): ?string
    {
        return $this->skype;
    }

    public function setSkype(?string $skype): self
    {
        $this->skype = $skype;

        return $this;
    }

    public function getZoomLink(): ?string
    {
        return $this->zoomLink;
    }

    public function setZoomLink(?string $zoomLink): self
    {
        $this->zoomLink = $zoomLink;

        return $this;
    }

    public function getResetPasswordToken(): ?string
    {
        return $this->resetPasswordToken;
    }

    public function setResetPasswordToken(?string $resetPasswordToken): self
    {
        $this->resetPasswordToken = $resetPasswordToken;

        return $this;
    }

    public function getEmailVerificationHash(): ?string
    {
        return $this->emailVerificationHash;
    }

    public function setEmailVerificationHash(?string $emailVerificationHash): self
    {
        $this->emailVerificationHash = $emailVerificationHash;

        return $this;
    }

    public function getPasswordResetHashExpiredAt(): ?\DateTimeImmutable
    {
        return $this->passwordResetHashExpiredAt;
    }

    public function setPasswordResetHashExpiredAt(?\DateTimeImmutable $passwordResetHashExpiredAt): self
    {
        $this->passwordResetHashExpiredAt = $passwordResetHashExpiredAt;

        return $this;
    }

    public function getIsEmailVerified(): ?bool
    {
        return $this->isEmailVerified;
    }

    public function setIsEmailVerified(?bool $isEmailVerified): self
    {
        $this->isEmailVerified = $isEmailVerified;

        return $this;
    }

    public function getTextPresentation(): ?string
    {
        return $this->textPresentation;
    }

    public function setTextPresentation(?string $textPresentation): self
    {
        $this->textPresentation = $textPresentation;

        return $this;
    }

    public function getVideoPresentation(): ?string
    {
        return $this->videoPresentation;
    }

    public function setVideoPresentation(?string $videoPresentation): self
    {
        $this->videoPresentation = $videoPresentation;

        return $this;
    }

    public function getPhoto(): ?string
    {
        return $this->photo;
    }

    public function setPhoto(?string $photo): self
    {
        $this->photo = $photo;

        return $this;
    }

    public function getCountryOfBirth(): ?string
    {
        return $this->countryOfBirth;
    }

    public function setCountryOfBirth(?string $countryOfBirth): self
    {
        $this->countryOfBirth = $countryOfBirth;

        return $this;
    }

    public function getCurrentCountry(): ?string
    {
        return $this->currentCountry;
    }

    public function setCurrentCountry(?string $currentCountry): self
    {
        $this->currentCountry = $currentCountry;

        return $this;
    }

    public function getSubjects(): ?array
    {
        return $this->subjects;
    }

    public function setSubjects(?array $subjects): self
    {
        $this->subjects = $subjects;

        return $this;
    }

    public function getBirthday(): ?\DateTimeInterface
    {
        return $this->birthday;
    }

    public function setBirthday(\DateTimeInterface $birthday): self
    {
        $this->birthday = $birthday;

        return $this;
    }

    public function getTimezone(): ?string
    {
        return $this->timezone;
    }

    public function setTimezone(string $timezone): self
    {
        $this->timezone = $timezone;

        return $this;
    }

    public function getGender(): ?string
    {
        return $this->gender;
    }

    public function setGender(string $gender): self
    {
        $this->gender = $gender;

        return $this;
    }

    public function getCreatedAt(): ?\DateTimeImmutable
    {
        return $this->createdAt;
    }

    public function setCreatedAt(\DateTimeImmutable $createdAt): self
    {
        $this->createdAt = $createdAt;

        return $this;
    }

    public function getUpdatedAt(): ?\DateTimeImmutable
    {
        return $this->updatedAt;
    }

    public function setUpdatedAt(\DateTimeImmutable $updatedAt): self
    {
        $this->updatedAt = $updatedAt;

        return $this;
    }

    public function getDeletedAt(): ?\DateTimeImmutable
    {
        return $this->deletedAt;
    }

    public function setDeletedAt(?\DateTimeImmutable $deletedAt): self
    {
        $this->deletedAt = $deletedAt;

        return $this;
    }


    /**
     * @return Collection<int, Lesson>
     */
    public function getLessons(): Collection
    {
        return $this->lessons;
    }

    public function addLesson(Lesson $lesson): self
    {
        if (!$this->lessons->contains($lesson)) {
            $this->lessons[] = $lesson;
            $lesson->setStudent($this);
        }

        return $this;
    }

    public function removeLesson(Lesson $lesson): self
    {
        if ($this->lessons->removeElement($lesson)) {
            // set the owning side to null (unless already changed)
            if ($lesson->getStudent() === $this) {
                $lesson->setStudent(null);
            }
        }

        return $this;
    }

    public function getTonRate(): ?float
    {
        return $this->tonRate;
    }

    public function setTonRate(?float $tonRate): self
    {
        $this->tonRate = $tonRate;

        return $this;
    }

    public function getTutorLanguages(): ?array
    {
        return $this->tutorLanguages;
    }

    public function setTutorLanguages(?array $tutorLanguages): self
    {
        $this->tutorLanguages = $tutorLanguages;

        return $this;
    }

    /**
     * @return Collection<int, UserBalance>
     */
    public function getUserBalances(): Collection
    {
        return $this->userBalances;
    }

    public function addUserBalance(UserBalance $userBalance): self
    {
        if (!$this->userBalances->contains($userBalance)) {
            $this->userBalances[] = $userBalance;
            $userBalance->setUser($this);
        }

        return $this;
    }

    public function removeUserBalance(UserBalance $userBalance): self
    {
        if ($this->userBalances->removeElement($userBalance)) {
            // set the owning side to null (unless already changed)
            if ($userBalance->getUser() === $this) {
                $userBalance->setUser(null);
            }
        }

        return $this;
    }

    public function getTonWalletAddress(): ?string
    {
        return $this->tonWalletAddress;
    }

    public function setTonWalletAddress(?string $tonWalletAddress): self
    {
        $this->tonWalletAddress = $tonWalletAddress;

        return $this;
    }

    /**
     * @return Collection<int, Review>
     */
    public function getReviews(): Collection
    {
        return $this->reviews;
    }

    public function addReview(Review $review): self
    {
        if (!$this->reviews->contains($review)) {
            $this->reviews[] = $review;
            $review->setUser($this);
        }

        return $this;
    }

    public function removeReview(Review $review): self
    {
        if ($this->reviews->removeElement($review)) {
            // set the owning side to null (unless already changed)
            if ($review->getUser() === $this) {
                $review->setUser(null);
            }
        }

        return $this;
    }

    public function getTimetableConfig(): ?TimetableConfig
    {
        return $this->timetableConfig;
    }

    public function setTimetableConfig(TimetableConfig $timetableConfig): self
    {
        // set the owning side of the relation if necessary
        if ($timetableConfig->getTutor() !== $this) {
            $timetableConfig->setTutor($this);
        }

        $this->timetableConfig = $timetableConfig;

        return $this;
    }

    /**
     * @return Collection<int, TimeslotReservation>
     */
    public function getTimeslotReservations(): Collection
    {
        return $this->timeslotReservations;
    }

    public function addTimeslotReservation(TimeslotReservation $timeslotReservation): self
    {
        if (!$this->timeslotReservations->contains($timeslotReservation)) {
            $this->timeslotReservations[] = $timeslotReservation;
            $timeslotReservation->setTutor($this);
        }

        return $this;
    }

    public function removeTimeslotReservation(TimeslotReservation $timeslotReservation): self
    {
        if ($this->timeslotReservations->removeElement($timeslotReservation)) {
            // set the owning side to null (unless already changed)
            if ($timeslotReservation->getTutor() === $this) {
                $timeslotReservation->setTutor(null);
            }
        }

        return $this;
    }

    /**
     * @return mixed
     */
    public function getIsTutorProfilePublished()
    {
        return $this->isTutorProfilePublished;
    }

    /**
     * @param mixed $isTutorProfilePublished
     */
    public function setIsTutorProfilePublished($isTutorProfilePublished): void
    {
        $this->isTutorProfilePublished = $isTutorProfilePublished;
    }
}
