export const BASE_URL = 'http://localhost:3000/';

export const jobOptions = {
	roles: ['Frontend', 'Backend', 'Fullstack'],
	levels: ['Junior', 'Mid', 'Senior'],
	contracts: ['Full Time', 'Part Time', 'Contract'],
	locations: ['US Only', 'UK Only', 'Worldwide', 'Remote', 'Other'],
	languages: ['JavaScript', 'Python', 'HTML', 'CSS', 'Ruby', 'Other'],
	tools: ['React', 'Vue', 'Sass', 'Django', 'RoR', 'Other']
};

export const initialFilters = {
	role: '',
	level: '',
	languages: [],
	tools: []
};

export const daysNew = 2