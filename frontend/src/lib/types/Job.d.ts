interface Job {
  jobId: number;
  company: string;
  logo: string;
  featured: boolean;
  position: string;
  role: string;
  level: string;
  contract: string;
  location: string;
  languages?: string[];
  tools?: string[];
  postedAt: string;
  isDeleted: boolean;
}

export default Job;
